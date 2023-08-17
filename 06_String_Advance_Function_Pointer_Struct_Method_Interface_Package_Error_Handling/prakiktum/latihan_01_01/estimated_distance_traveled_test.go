package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEstimatedDistanceTraveled(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		expected string
	}{
		{
			name: "SUV test",
			car : Car{
				carType: "SUV",
				fuelin:  10.5,
			},
			expected: "Car type: SUV, Est. distance: 15.75",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string = test.car.EstimatedDistanceTraveled()

			require.Equal(t, test.expected, result)
		})
	}
}