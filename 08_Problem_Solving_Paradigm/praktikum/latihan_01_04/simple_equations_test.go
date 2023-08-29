package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleEquations(t *testing.T) {
	tests := []struct{
		name string
		a, b, c int
		expected bool
	}{
		{
			name: "no solution",
			a: 1,
			b: 2,
			c: 3,
			expected: false,
		},
		{
			name: "6 6 14",
			a: 6,
			b: 6,
			c: 14,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result = simpleEquations(test.a, test.b, test.c)
			require.Equal(t, test.expected, result)
		})
	}
}