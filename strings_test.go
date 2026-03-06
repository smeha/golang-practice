package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		word string
		arr  []string
		want bool
	}{
		{"cat", []string{"c", "x", "a", "y", "t"}, true},
		{"cat", []string{"c", "a"}, false},
		{"", []string{"a", "b"}, true},
		{"abc", []string{}, false},
		{"café", []string{"c", "a", "f", "é"}, true},
	}
	for _, tt := range tests {
		if got := IsSubsequence(tt.word, tt.arr); got != tt.want {
			t.Errorf("IsSubsequence(%q, %v) = %v, want %v", tt.word, tt.arr, got, tt.want)
		}
	}
}

func TestLongestSubstringWithoutRepeating(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef", 6},
		{"café", 4},
	}
	for _, tt := range tests {
		if got := LongestSubstringWithoutRepeating(tt.s); got != tt.want {
			t.Errorf("LongestSubstringWithoutRepeating(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"café", "éfac"},
		{"abcd", "dcba"},
	}
	for _, tt := range tests {
		if got := ReverseString(tt.s); got != tt.want {
			t.Errorf("ReverseString(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}
