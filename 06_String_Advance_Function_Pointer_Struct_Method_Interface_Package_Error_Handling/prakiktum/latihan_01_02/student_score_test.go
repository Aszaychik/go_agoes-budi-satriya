package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var students Students = Students{
	{"Rizky", 80},
	{"Kobar", 75},
	{"Ismail", 70},
	{"Umam", 75},
	{"Yopan", 60},
}

func TestAverageStudentScore(t *testing.T) {
	tests := []struct {
		name     string
		students  Students
		expected int
	}{
		{
			name: "",
			students: students,
			expected: 72,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int = test.students.AverageStudentScore()
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSortStudentScore(t *testing.T) {
	tests := []struct {
		name     string
		students  Students
		expected Students
	}{
		{
			name: "Sort Student Score",
			students: students,
			expected: Students{
				{"Yopan", 60},
				{"Ismail", 70},
				{"Kobar", 75},
				{"Umam", 75},
				{"Rizky", 80},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result Students = test.students.SortStudentScore()
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMinStudentScore(t *testing.T) {
	tests := []struct {
		name     string
		students  Students
		expected Students
	}{
		{
			name: "Min Student Score",
			students: students,
			expected: Students{
				{"Yopan", 60},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var studentsScoreSorted Students = students.SortStudentScore()
			var result Students = Students{studentsScoreSorted[0]}
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMaxStudentScore(t *testing.T) {
	tests := []struct {
		name     string
		students  Students
		expected Students
	}{
		{
			name: "Max Student Score",
			students: students,
			expected: Students{
				{"Rizky", 80},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var studentsScoreSorted Students = students.SortStudentScore()
			var result Students = Students{studentsScoreSorted[len(studentsScoreSorted)-1]}
			require.Equal(t, test.expected, result)
		})
	}
}