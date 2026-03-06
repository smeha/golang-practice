package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{-1, false},
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{9, false},
		{17, true},
		{25, false},
		{97, true},
	}
	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{-5, -3, -8},
	}
	for _, tt := range tests {
		if got := Sum(tt.a, tt.b); got != tt.want {
			t.Errorf("Sum(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
