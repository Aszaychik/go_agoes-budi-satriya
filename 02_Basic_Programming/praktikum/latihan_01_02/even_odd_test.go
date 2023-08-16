package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected string
	}{
		{
			name:     "Even",
			number:   6,
			expected: "Even",
		},
		{
			name:     "Odd",
			number:   7,
			expected: "Odd",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string = EvenOrOdd(test.number)

			require.Equal(t, test.expected, result)
		})
	}
}