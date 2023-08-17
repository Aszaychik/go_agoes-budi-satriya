package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		a, b   string
		expected string
	}{
		{
			name: "AKA",
			a: "AKA",
			b: "AKASHI",
			expected: "AKA",
		},
		{
			name: "KANGOORO",
			a: "KANGOORO",
			b: "KANG",
			expected: "KANG",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string = Compare(test.a, test.b)

			require.Equal(t, test.expected, result)
		})
	}

}