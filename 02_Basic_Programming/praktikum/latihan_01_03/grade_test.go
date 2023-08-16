package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGradeCalculation(t *testing.T){
	tests := []struct {
		name     string
		nilai   int
		expected string
	}{
		{
			name:     "A",
			nilai:   90,
			expected: "A",
		},
		{
			name:     "B",
			nilai:   70,
			expected: "B",
		},
		{
			name:     "C",
			nilai:   60,
			expected: "C",
		},
		{
			name:     "D",
			nilai:   40,
			expected: "D",
		},
		{
			name:     "E",
			nilai:   20,
			expected: "E",
		},
		{
			name:     "Nilai Invalid",
			nilai:   101,
			expected: "Nilai Invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string = GradeCalculation(test.nilai)

			require.Equal(t, test.expected, result)
		})
	}
}