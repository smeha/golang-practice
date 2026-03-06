package main

import (
	"context"
	"testing"
	"time"
)

func TestTTLCacheSetGet(t *testing.T) {
	c := NewTTLCache()
	c.Set("key", "value", time.Second)

	v, ok := c.Get("key")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if v != "value" {
		t.Errorf("got %v, want %q", v, "value")
	}
}

func TestTTLCacheExpiry(t *testing.T) {
	c := NewTTLCache()
	c.Set("key", "value", 50*time.Millisecond)
	time.Sleep(100 * time.Millisecond)

	if _, ok := c.Get("key"); ok {
		t.Error("expected key to have expired")
	}
}

func TestTTLCacheMiss(t *testing.T) {
	c := NewTTLCache()
	if _, ok := c.Get("missing"); ok {
		t.Error("expected miss for non-existent key")
	}
}

func TestTTLCacheOverwrite(t *testing.T) {
	c := NewTTLCache()
	c.Set("key", "first", time.Second)
	c.Set("key", "second", time.Second)

	v, ok := c.Get("key")
	if !ok {
		t.Fatal("expected key to exist after overwrite")
	}
	if v != "second" {
		t.Errorf("got %v, want %q", v, "second")
	}
}

func TestTTLCacheJanitor(t *testing.T) {
	c := NewTTLCache()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c.StartJanitor(ctx, 50*time.Millisecond)

	c.Set("short", 1, 30*time.Millisecond)
	c.Set("long", 2, time.Second)

	time.Sleep(150 * time.Millisecond)

	c.mu.RLock()
	_, shortExists := c.items["short"]
	_, longExists := c.items["long"]
	c.mu.RUnlock()

	if shortExists {
		t.Error("janitor should have evicted key 'short'")
	}
	if !longExists {
		t.Error("janitor should not have evicted key 'long'")
	}
}
