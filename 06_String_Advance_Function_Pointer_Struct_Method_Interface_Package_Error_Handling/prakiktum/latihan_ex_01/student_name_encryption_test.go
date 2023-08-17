package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		student Student
		expected string
	}{
		{
			name: "rizky",
			student: Student{
				name:"rizky",
			},
			expected: "irapb",
		},
		{
			name: "agoes",
			student: Student{
				name:"agoes",
			},
			expected: "ztlvh",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string = test.student.Encode()

			require.Equal(t, test.expected, result)
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		student Student
		expected string
	}{
		{
			name: "rizky",
			student: Student{
				name:"irapb",
			},
			expected: "rizky",
		},
		{
			name: "agoes",
			student: Student{
				name:"ztlvh",
			},
			expected: "agoes",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string = test.student.Decode()

			require.Equal(t, test.expected, result)
		})
	}
}