package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPairSum(t *testing.T) {
	tests := []struct {
		name     string
		arr []int
		target int
		expected []int
	}{
		{
			name: "1",
			arr: []int{1, 2, 3, 4, 6},
			target: 6,
			expected: []int{1, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []int = PairSum(test.arr, test.target)

			require.Equal(t, test.expected, result)
		})
	}
}