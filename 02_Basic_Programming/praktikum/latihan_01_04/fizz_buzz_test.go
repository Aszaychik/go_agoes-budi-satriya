package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFizzBuzz(t *testing.T){
	tests := []struct {
		name string
		n  int
		expected []string
	}{

		{
			name: "15",
			n: 15,
			expected: []string{
				"1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []string = FizzBuzz(test.n)

			require.Equal(t, test.expected, result)
		})
	}
}