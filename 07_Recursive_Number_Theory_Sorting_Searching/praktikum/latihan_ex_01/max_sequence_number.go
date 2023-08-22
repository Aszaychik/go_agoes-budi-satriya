package main

import (
	"fmt"
	"math"
)


func MaxSequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var currentMax int
	var overallMax int
	
	for i := 1; i < len(arr); i++ {
		// Menghitung nilai baru currentMax dengan mengambil maksimum dari elemen saat ini dan jumlah elemen saat ini dan currentMax.
		currentMax = int(math.Max(float64(arr[i]), float64(currentMax+arr[i])))

		// Memperbarui overallMax dengan mengambil maksimum dari overallMax saat ini dan currentMax yang baru.
		overallMax = int(math.Max(float64(overallMax), float64(currentMax)))
	}

	return overallMax
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}