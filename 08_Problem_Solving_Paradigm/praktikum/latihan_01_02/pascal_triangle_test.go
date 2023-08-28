package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPascalTriangle(t *testing.T) {
	tests := []struct{
		name string
		numRows int
		expected [][]int
	}{
		{
			name: "5",
			numRows: 5,
			expected: [][]int{{1},{1,1},{1,2,1},{1,3,3,1},{1,4,6,4,1}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result [][]int = pascalTriangle(test.numRows)
			require.Equal(t, test.expected, result)
		})
	}
}