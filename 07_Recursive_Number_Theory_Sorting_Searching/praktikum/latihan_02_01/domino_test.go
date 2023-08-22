package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayinDomino(t *testing.T){
	tests := []struct {
		name     string
		cards   [][]int
		deck []int
		expected interface{}
	}{
		{
			name:     "[3,4]",
			cards:   [][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}},
			deck: []int{4, 3},
			expected: []int{3, 4},
		},
		{
			name:     "Tutup kartu",
			cards:   [][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}},
			deck: []int{5, 1},
			expected: "Tutup kartu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result interface{} = PlayingDomino(test.cards, test.deck)

			require.Equal(t, test.expected, result)
		})
	}
}