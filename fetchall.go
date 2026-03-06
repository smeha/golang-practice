package main

import (
	"context"
	"io"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup

	for i, u := range urls {
		i, u := i, u
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				<-sem
				wg.Done()
			}()

			r := FetchResult{URL: u}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				r.Err = err
				results[i] = r
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				r.Err = err
				results[i] = r
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
			results[i] = r
		}()
	}

	wg.Wait()
	return results
}
