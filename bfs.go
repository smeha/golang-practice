package main

import "sync"

type bfsResult struct {
	start int
	order []int
}

// ConcurrentBFSQueries runs BFS for each start node using a worker pool.
func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	if numWorkers <= 0 {
		numWorkers = 1
	}

	jobs := make(chan int)
	results := make(chan bfsResult)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for start := range jobs {
				order := bfs(graph, start)
				results <- bfsResult{start: start, order: order}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		for _, q := range queries {
			jobs <- q
		}
		close(jobs)
	}()

	out := make(map[int][]int, len(queries))
	for r := range results {
		out[r.start] = r.order
	}
	return out
}

func bfs(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	queue := make([]int, 0, 16)
	order := make([]int, 0, 16)

	visited[start] = true
	queue = append(queue, start)

	for head := 0; head < len(queue); head++ {
		u := queue[head]
		order = append(order, u)
		for _, v := range graph[u] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	return order
}
