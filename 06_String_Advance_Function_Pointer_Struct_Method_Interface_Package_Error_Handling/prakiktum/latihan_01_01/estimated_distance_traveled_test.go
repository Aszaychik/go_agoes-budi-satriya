package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEstimatedDistanceTraveled(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		expected float32
	}{
		{
			name: "SUV test",
			car : Car{
				carType: "SUV",
				fuelin:  10.5,
			},
			expected: 15.75,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result float32 = test.car.EstimatedDistanceTraveled()
			require.Equal(t, test.expected, result)
		})
	}
}