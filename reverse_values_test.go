package main

import "testing"

func TestReverseValues(t *testing.T) {
	tests := []struct {
		a, b        string
		wantA, wantB string
		wantErr     bool
	}{
		{"hello", "world", "world", "hello", false},
		{"", "world", "", "", true},
		{"hello", "", "", "", true},
		{"", "", "", "", true},
	}
	for _, tt := range tests {
		a, b, err := reverseValues(tt.a, tt.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("reverseValues(%q, %q) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && (a != tt.wantA || b != tt.wantB) {
			t.Errorf("reverseValues(%q, %q) = (%q, %q), want (%q, %q)", tt.a, tt.b, a, b, tt.wantA, tt.wantB)
		}
	}
}
