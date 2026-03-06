package main

import (
	"context"
	"time"
)

type Limiter struct {
	tokens chan struct{}
	stop   chan struct{}
	ticker *time.Ticker
}

func NewLimiter(rps int, burst int) *Limiter {
	if rps <= 0 {
		rps = 1
	}
	if burst <= 0 {
		burst = 1
	}

	l := &Limiter{
		tokens: make(chan struct{}, burst),
		stop:   make(chan struct{}),
		ticker: time.NewTicker(time.Second / time.Duration(rps)),
	}
	for i := 0; i < burst; i++ {
		l.tokens <- struct{}{}
	}

	go func() {
		defer l.ticker.Stop()
		for {
			select {
			case <-l.stop:
				return
			case <-l.ticker.C:
				select {
				case l.tokens <- struct{}{}:
				default:
					// bucket full
				}
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

func (l *Limiter) Stop() {
	close(l.stop)
}
