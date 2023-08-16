package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestAreaTrapezoid(t *testing.T){
	tests := []struct{
		name string
		base1, base2, height, expected float32
	}{
		{
			name: "1",
			base1: 10,
			base2: 20,
			height: 30,
			expected: 450,
		},
		{
			name: "2",
			base1: 4,
			base2: 6,
			height: 5,
			expected: 25,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result float32 = AreaOfTrapezoid(test.base1, test.base2, test.height)

			require.Equal(t, test.expected, result)
		})
	}
}