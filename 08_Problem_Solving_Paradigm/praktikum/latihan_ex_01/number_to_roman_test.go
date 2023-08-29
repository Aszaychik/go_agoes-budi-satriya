package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToRomawi(t *testing.T){
	tests := []struct{
		name string
		number int
		expected string
	}{
		{
			name: "Exceeds",
			number: 4000,
			expected: "Exceeds the maximum Roman number",
		},
		{
			name: "23",
			number: 23,
			expected: "XXIII",
		},
		{
			name: "2021",
			number: 2021,
			expected: "MMXXI",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result = toRomawi(test.number)
			require.Equal(t, test.expected, result)
		})
	}
}