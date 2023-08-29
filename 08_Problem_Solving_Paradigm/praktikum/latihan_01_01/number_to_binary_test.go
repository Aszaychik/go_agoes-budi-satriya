package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToBinary(t *testing.T){
	tests := []struct{
		name string
		number int
		expected []string
	}{
		{
			name: "3",
			number: 3,
			expected: []string{"0", "1", "10", "11"},
		},
		{
			name: "10",
			number: 10,
			expected: []string{"0","1","10","11","100","101","110","111","1000","1001","1010"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []string = toBinary(test.number)
			require.Equal(t, test.expected, result)
		})
	}
}