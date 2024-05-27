package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "fizz"},
		{5, "buzz"},
		{15, "fizzbuzz"},
		{30, "fizzbuzz"},
		{7, "7"},
	}

	for _, test := range tests {
		result := fizzBuzz(test.input)
		if result != test.expected {
			t.Errorf("fizzBuzz(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
