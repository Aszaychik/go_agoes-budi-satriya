package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMinMax(t *testing.T) {
	tests := []struct {
		name     string
		numbers     []int
		expectedMin, expectedMax int
	}{
		{
			name:     "",
			numbers: []int{1, 2, 3, 9, 7, 8},
			expectedMax: 9,
			expectedMin: 1,
		},
		{
			name:     "",
			numbers: []int{100, 200, 300, 900, 700, 800},
			expectedMax: 900,
			expectedMin: 100,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var request []*int = make([]*int, len(test.numbers))
			for i, n := range test.numbers {
				var number int = n // Create a new variable to hold value of n to avoid overwriting the original n
				request[i] = &number
			}
			
			var min, max int = GetMinMax(request ...)

			require.Equal(t, test.expectedMin, min)
			require.Equal(t, test.expectedMax, max)
		})
	}
}