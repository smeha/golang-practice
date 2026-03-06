package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Available exercises:")
	fmt.Println("  bfs            - Concurrent BFS queries on a graph")
	fmt.Println("  manager        - Employee manager (add/remove/find/avg salary)")
	fmt.Println("  longestrepeat  - Longest substring without repeating characters")
	fmt.Println("  subseq         - Check if word is a subsequence of an array")
	fmt.Println("  sum            - Sum two integers")
	fmt.Println("  prime          - Check if a number is prime")
	fmt.Println("  reverse        - Reverse a string")
	fmt.Println("  reversevalues  - Swap two input values")
	fmt.Println("  limiter        - Token bucket rate limiter demo")
	fmt.Println("  ttlcache       - TTL cache demo")
	fmt.Println("  fetchall       - Concurrent URL fetcher demo")
	fmt.Println()
	fmt.Print("Enter exercise name: ")

	fn, _ := reader.ReadString('\n')
	fn = strings.TrimSpace(fn)

	switch fn {
	case "bfs":
		graph := map[int][]int{
			0: {1, 2},
			1: {2, 3},
			2: {3},
			3: {4},
			4: {},
		}
		queries := []int{0, 1, 2}
		results := ConcurrentBFSQueries(graph, queries, 2)
		for k, v := range results {
			fmt.Printf("  BFS from %d: %v\n", k, v)
		}

	case "manager":
		m := Manager{}
		m.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
		m.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
		m.AddEmployee(Employee{ID: 3, Name: "Carol", Age: 27, Salary: 90000})
		m.RemoveEmployee(1)
		fmt.Printf("  Average salary (after removing Alice): %.2f\n", m.GetAverageSalary())
		if e := m.FindEmployeeByID(2); e != nil {
			fmt.Printf("  Found: %+v\n", *e)
		}

	case "longestrepeat":
		fmt.Print("  Enter string: ")
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		fmt.Printf("  Longest substring without repeating: %d\n", LongestSubstringWithoutRepeating(s))

	case "subseq":
		arr := []string{"c", "x", "a", "y", "t"}
		fmt.Printf("  Array: %v\n", arr)
		fmt.Print("  Enter word: ")
		w, _ := reader.ReadString('\n')
		w = strings.TrimSpace(w)
		fmt.Printf("  Is subsequence: %v\n", IsSubsequence(w, arr))

	case "sum":
		fmt.Print("  Enter first number: ")
		aStr, _ := reader.ReadString('\n')
		fmt.Print("  Enter second number: ")
		bStr, _ := reader.ReadString('\n')
		a, _ := strconv.Atoi(strings.TrimSpace(aStr))
		b, _ := strconv.Atoi(strings.TrimSpace(bStr))
		fmt.Printf("  Sum: %d\n", Sum(a, b))

	case "prime":
		fmt.Print("  Enter number: ")
		numStr, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(numStr))
		fmt.Printf("  %d is prime: %v\n", n, IsPrime(n))

	case "reverse":
		fmt.Print("  Enter string: ")
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		fmt.Printf("  Reversed: %s\n", ReverseString(s))

	case "reversevalues":
		fmt.Print("  Enter first value: ")
		a, _ := reader.ReadString('\n')
		fmt.Print("  Enter second value: ")
		b, _ := reader.ReadString('\n')
		v1, v2, err := reverseValues(strings.TrimSpace(a), strings.TrimSpace(b))
		if err != nil {
			fmt.Println("  Error:", err)
		} else {
			fmt.Printf("  Swapped: %q %q\n", v1, v2)
		}

	case "limiter":
		fmt.Println("  Creating limiter: 5 rps, burst 3. Acquiring 6 tokens...")
		l := NewLimiter(5, 3)
		defer l.Stop()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		for i := 1; i <= 6; i++ {
			start := time.Now()
			if err := l.Acquire(ctx); err != nil {
				fmt.Printf("  Token %d: error - %v\n", i, err)
				break
			}
			fmt.Printf("  Token %d acquired after %v\n", i, time.Since(start).Round(time.Millisecond))
		}

	case "ttlcache":
		cache := NewTTLCache()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cache.StartJanitor(ctx, 500*time.Millisecond)

		cache.Set("hello", "world", 200*time.Millisecond)
		cache.Set("foo", "bar", 2*time.Second)

		if v, ok := cache.Get("hello"); ok {
			fmt.Printf("  hello = %v\n", v)
		}
		fmt.Println("  Waiting 300ms for 'hello' to expire...")
		time.Sleep(300 * time.Millisecond)
		if _, ok := cache.Get("hello"); !ok {
			fmt.Println("  hello expired (not found)")
		}
		if v, ok := cache.Get("foo"); ok {
			fmt.Printf("  foo = %v (still alive)\n", v)
		}

	case "fetchall":
		urls := []string{
			"https://example.com",
			"https://httpbin.org/status/404",
			"https://httpbin.org/delay/5", // will timeout
		}
		fmt.Printf("  Fetching %d URLs with 2s timeout...\n", len(urls))
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		client := &http.Client{Timeout: 1500 * time.Millisecond}
		results := FetchAll(ctx, client, urls, 3)
		for _, r := range results {
			errStr := "nil"
			if r.Err != nil {
				errStr = r.Err.Error()
			}
			fmt.Printf("  %s  status=%d  err=%s\n", r.URL, r.Status, errStr)
		}

	default:
		fmt.Println("Unknown exercise.")
	}
}
