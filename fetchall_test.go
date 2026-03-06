package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestFetchAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("hello"))
		case "/notfound":
			w.WriteHeader(http.StatusNotFound)
		case "/slow":
			time.Sleep(500 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer srv.Close()

	urls := []string{
		srv.URL + "/ok",
		srv.URL + "/notfound",
		srv.URL + "/slow",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	results := FetchAll(ctx, &http.Client{}, urls, 3)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
	if results[0].Status != http.StatusOK {
		t.Errorf("results[0] status = %d, want 200", results[0].Status)
	}
	if string(results[0].Body) != "hello" {
		t.Errorf("results[0] body = %q, want %q", string(results[0].Body), "hello")
	}
	if results[1].Status != http.StatusNotFound {
		t.Errorf("results[1] status = %d, want 404", results[1].Status)
	}
	if results[2].Err == nil {
		t.Error("expected error for slow request due to timeout")
	}
	if results[2].URL != urls[2] {
		t.Errorf("results[2] URL = %q, want %q", results[2].URL, urls[2])
	}
}

func TestFetchAllPreservesOrder(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Introduce small delay so goroutines finish out of order
		if r.URL.Path == "/first" {
			time.Sleep(30 * time.Millisecond)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(r.URL.Path))
	}))
	defer srv.Close()

	urls := []string{srv.URL + "/first", srv.URL + "/second"}
	results := FetchAll(context.Background(), &http.Client{}, urls, 2)

	if string(results[0].Body) != "/first" {
		t.Errorf("results[0] body = %q, want %q", string(results[0].Body), "/first")
	}
	if string(results[1].Body) != "/second" {
		t.Errorf("results[1] body = %q, want %q", string(results[1].Body), "/second")
	}
}

func TestFetchAllConcurrencyLimit(t *testing.T) {
	var mu sync.Mutex
	var active, maxActive int

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		active++
		if active > maxActive {
			maxActive = active
		}
		mu.Unlock()

		time.Sleep(20 * time.Millisecond)

		mu.Lock()
		active--
		mu.Unlock()

		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	urls := make([]string, 6)
	for i := range urls {
		urls[i] = srv.URL + "/test"
	}

	FetchAll(context.Background(), &http.Client{}, urls, 2)

	if maxActive > 2 {
		t.Errorf("max concurrent requests = %d, want <= 2", maxActive)
	}
}

func TestFetchAllEmpty(t *testing.T) {
	results := FetchAll(context.Background(), &http.Client{}, nil, 5)
	if len(results) != 0 {
		t.Errorf("expected 0 results, got %d", len(results))
	}
}
