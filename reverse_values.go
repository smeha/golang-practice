package main

import "fmt"

func reverseValues(a, b string) (string, string, error) {
	if a == "" || b == "" {
		return "", "", fmt.Errorf("inputs must not be empty")
	}
	return b, a, nil
}
