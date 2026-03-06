package main

import (
	"context"
	"time"
)

type Limiter struct {
	tokens chan struct{}
}

func NewLimiter(rps int, burst int) *Limiter {
	if rps <= 0 {
		rps = 1
	}
	if burst <= 0 {
		burst = 1
	}

	l := &Limiter{tokens: make(chan struct{}, burst)}
	for i := 0; i < burst; i++ {
		l.tokens <- struct{}{}
	}

	interval := time.Second / time.Duration(rps)
	t := time.NewTicker(interval)

	go func() {
		defer t.Stop()
		for range t.C {
			select {
			case l.tokens <- struct{}{}:
			default:
				// bucket full
			}
		}
	}()

	return l
}

func (l *Limiter) Acquire(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-l.tokens:
		return nil
	}
}
