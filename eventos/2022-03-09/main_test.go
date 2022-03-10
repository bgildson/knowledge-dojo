package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  bool
	}{
		{
			name: "teste01",
			in:   121,
			out:  true,
		},
		{
			name: "teste02",
			in:   -121,
			out:  false,
		},
		{
			name: "teste03",
			in:   10,
			out:  false,
		},
		{
			name: "teste04",
			in:   1,
			out:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsPalindrome(test.in)
			if got != test.out {
				t.Errorf("want %v got %v", test.out, got)
			}
		})
	}
}
