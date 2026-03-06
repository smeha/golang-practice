package main

import (
	"context"
	"testing"
	"time"
)

func TestLimiterBurstAvailableImmediately(t *testing.T) {
	l := NewLimiter(1, 3)
	defer l.Stop()

	ctx := context.Background()
	start := time.Now()
	for i := 0; i < 3; i++ {
		if err := l.Acquire(ctx); err != nil {
			t.Fatalf("unexpected error on token %d: %v", i+1, err)
		}
	}
	if elapsed := time.Since(start); elapsed > 50*time.Millisecond {
		t.Errorf("burst tokens took too long: %v", elapsed)
	}
}

func TestLimiterContextCancellation(t *testing.T) {
	l := NewLimiter(1, 1)
	defer l.Stop()

	ctx := context.Background()
	_ = l.Acquire(ctx) // drain burst

	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	if err := l.Acquire(ctx); err == nil {
		t.Error("expected error when context times out waiting for token")
	}
}

func TestLimiterRateLimit(t *testing.T) {
	l := NewLimiter(10, 1) // 10 rps = 1 token per 100ms, burst 1
	defer l.Stop()

	ctx := context.Background()
	_ = l.Acquire(ctx) // drain burst

	start := time.Now()
	_ = l.Acquire(ctx) // must wait ~100ms for next token
	elapsed := time.Since(start)

	if elapsed < 50*time.Millisecond {
		t.Errorf("rate limiting not enforced: got token in %v, expected ~100ms", elapsed)
	}
}

func TestLimiterStop(t *testing.T) {
	l := NewLimiter(10, 1)
	l.Stop() // should not panic
}
