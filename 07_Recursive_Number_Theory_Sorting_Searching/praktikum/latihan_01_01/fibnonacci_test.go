package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	tests := []struct{
		name string
		number int
		expected int
	}{
		{
			name: "21",
			number: 21,
			expected: 10946,
		},
		{
			name: "12",
			number: 12,
			expected: 144,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int

			result = Fibonacci(test.number)

			require.Equal(t, test.expected, result)
		})
	}
}