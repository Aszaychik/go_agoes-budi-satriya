package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// Buat variabel firstIndex dan lastIndex untuk menyimpan index awal dan akhir dari slice
	var firstIndex int = 0
	var lastIndex int = len(arr) - 1

	// Iterasi selama `firstIndex` kurang dari `lastIndex`
	for firstIndex < lastIndex {
		// Hitung jumlah elemen pada index `firstIndex` dan `lastIndex`
		var currentSum int = arr[firstIndex] + arr[lastIndex]

		// Jika jumlah elemen tersebut sama dengan target, kembalikan slice yang berisi index `firstIndex` dan `lastIndex`
		if currentSum == target {
			return []int{firstIndex, lastIndex}
		} else if currentSum < target {
			// Increment `firstIndex`
			firstIndex++
		} else {
			// Decrement `lastIndex`
			lastIndex--
		}
	}

	// Jika tidak ada elemen yang memiliki jumlah sama dengan target, kembalikan `nil`
	return nil
}

func main() {

	// Test cases

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}