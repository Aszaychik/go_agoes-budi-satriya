package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumBuyProduct(t *testing.T){
	tests := []struct {
		name     string
		money   int
		productPrice []int
		expected int
	}{
		{
			name:     "",
			money:   30000,
			productPrice: []int{15000, 10000, 12000, 5000, 3000},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int = MaximumBuyProduct(test.money, test.productPrice)

			require.Equal(t, test.expected, result)
		})
	}
}