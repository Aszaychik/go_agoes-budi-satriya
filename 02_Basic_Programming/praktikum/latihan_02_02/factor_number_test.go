package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactorNumber(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected []int
	}{
		{
			name: "26",
			number: 26,
			expected:[]int{1,2,13,26}                                                                     ,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []int = FactorNumber(test.number)

			require.Equal(t, test.expected, result)
		})
	}
}