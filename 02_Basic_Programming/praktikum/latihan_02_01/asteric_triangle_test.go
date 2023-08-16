package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAstericTriangle(t *testing.T) {
	tests := []struct {
		name     string
		rows   int
		expected string
	}{
		{
			name: "5 star",
			rows: 5,
			expected:"    * \n   * * \n  * * * \n * * * * \n* * * * * \n"                                                                     ,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string = AstericTriangle(test.rows)

			require.Equal(t, test.expected, result)
		})
	}
}