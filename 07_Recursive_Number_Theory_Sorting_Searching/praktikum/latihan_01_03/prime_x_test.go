package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected bool
	}{
		{
			name:     "True",
			number:   1000000007,
			expected: true,
		},
		{
			name:     "False",
			number:   20,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result bool = isPrime(test.number)

			require.Equal(t, test.expected, result)
		})
	}
}

func TestPrimeX(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected int
	}{
		{
			name:     "29",
			number:   10,
			expected: 29,
		},
		{
			name:     "100",
			number:   100,
			expected: 541,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int = PrimeX(test.number)

			require.Equal(t, test.expected, result)
		})
	}
}