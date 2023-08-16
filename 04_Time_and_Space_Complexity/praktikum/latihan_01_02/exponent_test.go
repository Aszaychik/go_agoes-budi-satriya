package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPow(t *testing.T) {
	tests := []struct {
		name     string
		x, n   int
		expected int
	}{
		{
			name: "1",
			x: 7,
			n: 3,
			expected: 343,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int = Pow(test.x, test.n)

			require.Equal(t, test.expected, result)
		})
	}
}