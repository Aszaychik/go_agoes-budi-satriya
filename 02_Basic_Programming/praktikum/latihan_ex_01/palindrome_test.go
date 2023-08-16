package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrom(t *testing.T) {
	tests := []struct {
		name     string
		words   string
		expected bool
	}{
		{
			name: "aya aya",
			words: "aya aya",
			expected: true                                                                     ,
		},
		{
			name: "asz asz",
			words: "asz asz",
			expected: false                                                                     ,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result bool = IsPalindrome(test.words)

			require.Equal(t, test.expected, result)
		})
	}
}