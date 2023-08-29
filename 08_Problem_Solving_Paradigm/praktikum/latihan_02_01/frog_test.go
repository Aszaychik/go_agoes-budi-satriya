package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrog(t *testing.T){
	tests := []struct{
		name string
		jumps []int
		expected int
	}{
		{
			name: "30",
			jumps: []int{10, 30, 40, 20},
			expected: 30,
		},
		{
			name: "40",
			jumps: []int{30, 10, 60, 10, 60, 50},
			expected: 40,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result = frog(test.jumps)
			require.Equal(t, test.expected, result)
		})
	}
}