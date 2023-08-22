package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxSequence(t *testing.T){
	tests := []struct {
		name     string
		arr   []int
		expected int
		}{
			{
			name:     "",
			arr: []int{-2, -5, 6, 2, -3, 1, 6, -6},
			expected: 12,
			},
		}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int = MaxSequence(test.arr)

			require.Equal(t, test.expected, result)
		})
	}
}