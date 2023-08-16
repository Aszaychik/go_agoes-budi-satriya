package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrix2dSum(t *testing.T) {
	tests := []struct {
		name     string
		matrix [][]int
		expected int
	}{
		{
			name: "1",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{9, 8, 9},
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int = Matrix2dSum(test.matrix)

			require.Equal(t, test.expected, result)
		})
	}
}