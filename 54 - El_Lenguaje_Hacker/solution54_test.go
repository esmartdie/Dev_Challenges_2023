package main

import (
	"testing"
)

func TestHackerProcessing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "#3110 \\/\\/0r1)"},
		{"leet speak", "1337 5|*34>|"},
		{"Golang 123", "6014^/6 LRE"},
	}

	for _, test := range tests {
		result := hackerProcessing(test.input)
		if result != test.expected {
			t.Errorf("hackerProcessing(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
