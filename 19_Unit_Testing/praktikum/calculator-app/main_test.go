package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
  tests := []struct {
    name string
    firstNumber, secondNumber float64
    expected float64
  }{
    {
      name: "Test 11 + 10 = 21",
      firstNumber: 11,
      secondNumber: 10,
      expected: 21,
    },
    {
      name: "Test 111001 + 20011011 = 20122012",
      firstNumber: 111001,
      secondNumber: 20011011,
      expected: 20122012,
    },
    {
      name: "Test 111001 + 0 = 111001",
      firstNumber: 111001,
      secondNumber: 0,
      expected: 111001,
    },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T) {
      result := Addition(test.firstNumber, test.secondNumber)
			
			require.Equal(t, test.expected, result)
    })
  }
}

func TestSubtraction(t *testing.T) {
	tests := []struct {
		name     string
		firstNumber, secondNumber   float64
		expected float64
	}{
		{
			name: "Test 11 - 10 = 1",
			firstNumber: 11,
			secondNumber: 10,
			expected: 1,
		},
		{
			name: "Test 20011011 - 111001 = 19900010",
			firstNumber: 20011011,
			secondNumber: 111001,
			expected: 19900010,
		},
		{
			name: "Test 111001 - 20011011 = -19900010",
			firstNumber: 111001,
			secondNumber: 20011011,
			expected: -19900010,
		},
		{
			name: "Test 111001 - 0 = 111001",
			firstNumber: 111001,
			secondNumber: 0,
			expected: 111001,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Subtraction(test.firstNumber, test.secondNumber)

			require.Equal(t, test.expected, result)
		})
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		name     string
		firstNumber, secondNumber   float64
		expected float64
	}{
		{
			name: "Test 1000 / 10 = 100",
			firstNumber: 1000,
			secondNumber: 10,
			expected: 100,
		},
		{
			name: "Test 111001 / 20011011 = 0.0055469961012964315",
			firstNumber: 111001,
			secondNumber: 20011011,
			expected: 0.0055469961012964315,
		},
		{
			name: "Test 20011011 / 111001 = 180.27775425446617",
			firstNumber: 20011011,
			secondNumber: 111001,
			expected: 180.27775425446617,
		},
		{
			name: "Test 374985738495 / 0 = Cannot divide by zero",
			firstNumber: 374985738495,
			secondNumber: 0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Division(test.firstNumber, test.secondNumber)
			if err != nil {
				require.EqualError(t, err, "Cannot divide by zero")
			}
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		firstNumber, secondNumber   float64
		expected float64
	}{
		{
			name: "Test 11 * 10 = 110",
			firstNumber: 11,
			secondNumber: 10,
			expected: 110,
		},
		{
			name: "Test 111001 * 20011011 = 2221242232011",
			firstNumber: 111001,
			secondNumber: 20011011,
			expected: 2221242232011,
		},
		{
			name: "Test 111001 * 0 = 0",
			firstNumber: 111001,
			secondNumber: 0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Multiplication(test.firstNumber, test.secondNumber)

			require.Equal(t, test.expected, result)
		})
	}
}
