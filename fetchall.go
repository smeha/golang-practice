package main

import (
	"context"
	"io"
	"net/http"
)

type FetchResult struct {
	URL    string
	Body   []byte
	Status int
	Err    error
}

func FetchAll(ctx context.Context, client *http.Client, urls []string, maxConcurrent int) []FetchResult {
	if maxConcurrent <= 0 {
		maxConcurrent = 1
	}

	results := make([]FetchResult, len(urls))
	sem := make(chan struct{}, maxConcurrent)
	done := make(chan struct{})

	type item struct {
		i int
		r FetchResult
	}
	out := make(chan item)

	go func() {
		defer close(done)
		for it := range out {
			results[it.i] = it.r
		}
	}()

	for i, u := range urls {
		i, u := i, u
		sem <- struct{}{}
		go func() {
			defer func() { <-sem }()

			r := FetchResult{URL: u}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				r.Err = err
				select {
				case out <- item{i: i, r: r}:
				case <-ctx.Done():
				}
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				r.Err = err
				select {
				case out <- item{i: i, r: r}:
				case <-ctx.Done():
				}
				return
			}
			defer resp.Body.Close()

			r.Status = resp.StatusCode
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				r.Err = err
			} else {
				r.Body = body
			}

			select {
			case out <- item{i: i, r: r}:
			case <-ctx.Done():
			}
		}()
	}

	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
	close(out)
	<-done
	return results
}
