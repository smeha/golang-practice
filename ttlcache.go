package main

import (
	"context"
	"sync"
	"time"
)

type cacheItem struct {
	v       any
	expires time.Time
}

type TTLCache struct {
	mu    sync.RWMutex
	items map[string]cacheItem
}

func NewTTLCache() *TTLCache {
	return &TTLCache{items: make(map[string]cacheItem)}
}

func (c *TTLCache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = cacheItem{
		v:       value,
		expires: time.Now().Add(ttl),
	}
}

func (c *TTLCache) Get(key string) (any, bool) {
	c.mu.RLock()
	it, ok := c.items[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	if time.Now().After(it.expires) {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		return nil, false
	}
	return it.v, true
}

func (c *TTLCache) StartJanitor(ctx context.Context, interval time.Duration) {
	t := time.NewTicker(interval)
	go func() {
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				now := time.Now()
				c.mu.Lock()
				for k, it := range c.items {
					if now.After(it.expires) {
						delete(c.items, k)
					}
				}
				c.mu.Unlock()
			}
		}
	}()
}
