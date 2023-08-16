package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapping(t *testing.T) {
	tests := []struct {
		name     string
		slice []string
		expected map[string]int
	}{
		{
			name:     "1",
			slice:   []string{"asd", "qwe", "asd", "adi", "qwe", "qwe"},
			expected: map[string]int{"adi": 1, "asd": 2, "qwe": 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result map[string]int = Mapping(test.slice)

			require.Equal(t, test.expected, result)
		})
	}
}