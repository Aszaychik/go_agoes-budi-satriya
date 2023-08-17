package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCaesar(t *testing.T) {
	tests := []struct {
		name     string
		offset int
		input     string
		expected string
	}{
		{
			name: "Lower Case",
			offset: 1,
			input: "abcdefghijklmnopqrstuvwxyz",
			expected: "bcdefghijklmnopqrstuvwxyza",
		},
		{
			name: "Upper Case",
			offset: 1,
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expected: "BCDEFGHIJKLMNOPQRSTUVWXYZA",
		},
		{
			name: "Digit",
			offset: 2,
			input: "123",
			expected: "345",
		},
		{
			name: "N/A",
			offset: 2,
			input: "!@#$%^&*()",
			expected: "!@#$%^&*()",
		},
		{
			name: "Random",
			offset: 1,
			input: "J3x9j9!LtS",
			expected: "K4y0k0!MuT",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string = Caesar(test.offset, test.input)

			require.Equal(t, test.expected, result)
		})
	}
}