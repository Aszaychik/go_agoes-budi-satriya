package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMostAppearItem(t *testing.T) {
	tests := []struct{
		name string
		items []string
		expected []pair
	}{
		{
			name: "Programming Lang",
			items: []string{"js", "js", "golang", "ruby", "ruby", "js", "js"},
			expected: []pair{{"golang", 1}, {"ruby", 2}, {"js", 4}},
		},
		{
			name: "Alphabet",
			items: []string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"},
			expected: []pair{{"C", 1}, {"D", 2}, {"B", 3}, {"A", 4}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []pair = MostAppearItem(test.items)

			require.Equal(t, test.expected, result)
		})
	}
}