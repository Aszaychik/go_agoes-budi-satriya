package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayMerge(t *testing.T) {
	tests := []struct {
		name     string
		arrayA, arrayB []string
		expected []string
	}{
		{
			name:     "All Jin",
			arrayA:   []string{"jin", "jin", "jin"},
			arrayB: []string{"jin", "jin", "jin"},
			expected: []string{"jin"},
		},
		{
			name:     "All Unique",
			arrayA:   []string{"king", "devil jin", "akuma"},
			arrayB: []string{"eddie", "steve", "geese"},
			expected: []string{"king", "devil jin", "akuma", "eddie", "steve", "geese"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []string = ArrayMerge(test.arrayA, test.arrayB)

			require.Equal(t, test.expected, result)
		})
	}
}