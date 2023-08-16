package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrimeNumber(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected bool
	}{
		{
			name:     "Prime",
			number:   1000000007,
			expected: true,
		},
		{
			name:     "Not Prime",
			number:   35,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result bool = PrimeNumber(test.number)

			require.Equal(t, test.expected, result)
		})
	}
}