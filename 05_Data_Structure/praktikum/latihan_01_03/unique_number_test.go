package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueNumber(t *testing.T) {
	tests := []struct {
		name     string
		number    string
		expected []int
	}{
		{
			name: "1",
			number: "1234123",
			expected: []int{4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []int = UniqueNumber(test.number)

			require.Equal(t, test.expected, result)
		})
	}
}