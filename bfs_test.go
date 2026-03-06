package main

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}
	order := bfs(graph, 0)
	if order[0] != 0 {
		t.Errorf("first node should be start node 0, got %d", order[0])
	}
	pos := func(n int) int {
		for i, v := range order {
			if v == n {
				return i
			}
		}
		return -1
	}
	// 1 and 2 must appear before 3
	if pos(1) >= pos(3) || pos(2) >= pos(3) {
		t.Errorf("BFS order incorrect: %v", order)
	}
}

func TestBFSDisconnected(t *testing.T) {
	graph := map[int][]int{
		0: {1},
		2: {3},
	}
	order := bfs(graph, 0)
	if !reflect.DeepEqual(order, []int{0, 1}) {
		t.Errorf("expected [0 1], got %v", order)
	}
}

func TestBFSSingleNode(t *testing.T) {
	graph := map[int][]int{0: {}}
	order := bfs(graph, 0)
	if !reflect.DeepEqual(order, []int{0}) {
		t.Errorf("expected [0], got %v", order)
	}
}

func TestConcurrentBFSQueries(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {3},
		3: {4},
		4: {},
	}
	queries := []int{0, 1, 2}
	results := ConcurrentBFSQueries(graph, queries, 2)

	if len(results) != len(queries) {
		t.Fatalf("expected %d results, got %d", len(queries), len(results))
	}
	for _, q := range queries {
		r, ok := results[q]
		if !ok {
			t.Errorf("missing result for query start=%d", q)
			continue
		}
		if r[0] != q {
			t.Errorf("BFS from %d: first element should be %d, got %d", q, q, r[0])
		}
	}
}

func TestConcurrentBFSQueriesZeroWorkers(t *testing.T) {
	graph := map[int][]int{0: {1}, 1: {}}
	results := ConcurrentBFSQueries(graph, []int{0}, 0)
	if _, ok := results[0]; !ok {
		t.Error("expected result for query 0 even with 0 workers")
	}
}
