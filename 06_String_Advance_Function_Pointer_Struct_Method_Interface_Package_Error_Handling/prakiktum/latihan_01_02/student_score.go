package main

import (
	"fmt"
	"sort"
)

type Student struct{
	Name  string
	Score int
}

type Students []Student

func (s Students) AverageStudentScore() int {
	var result, totalScore int

	// Penggunaan for range untuk mengambil data dari slice dan menghitung total score
	for _, student := range s {
		totalScore += student.Score
	}

	// Menghitung rata-rata
	result = totalScore / len(s)

	return result
}

func (s Students) SortStudentScore() Students {
	// Mengcopy slice yg diberikan supaya tidak terjadi perubahan di slice original
	var sorted Students = s

	// Mengurutkan score dari yang terkecil ke yang terbesar
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Score < sorted[j].Score
	})

	return sorted
}

func main() {
	var students Students = Students{
		{"Rizky", 80},
		{"Kobar", 75},
		{"Ismail", 70},
		{"Umam", 75},
		{"Yopan", 60},
	}

	var studentsScoreSorted Students = students.SortStudentScore() // Mengambil score dari slice yg di urutkan
	
	var averageStudentScore int = students.AverageStudentScore() // Mengambil rata-rata score
	var minScoreStudent Students = Students{studentsScoreSorted[0]} // Mengambil score dari slice yg terkecil / index pertama
	var maxScoreStudent Students = Students{studentsScoreSorted[len(studentsScoreSorted)-1]} // Mengambil score dari slice yg terbesar / index terakhir

	fmt.Println("Average Score:", averageStudentScore)
	fmt.Println("Min Score of Students:",minScoreStudent)
	fmt.Println("Max Score of Students:",maxScoreStudent)
}