package main

import "fmt"

func UniqueNumber(number string) []int {
	// Buat slice kosong untuk menyimpan hasil
	var result []int

	// Buat map untuk menyimpan jumlah kemunculan masing-masing digit
	var counts map[int32]int = make(map[int32]int)

	// Iterasi setiap digit dalam string
	for _, digit := range number {
		// Increment jumlah kemunculan digit tersebut
		counts[digit]++
	}

	// Iterasi setiap digit dalam string lagi
	for _, digit := range number {
		// Cek apakah digit tersebut muncul satu kali
		if counts[digit] == 1 {
			// Tambahkan digit tersebut ke slice hasil
			result = append(result, int(digit-'0'))

			// Set jumlah kemunculan digit tersebut menjadi -1 untuk menandakan bahwa digit tersebut sudah diproses
			counts[digit] = -1
		}
	}

	// Kembalikan slice hasil
	return result
}


func main() {

	// Test cases

	fmt.Println(UniqueNumber("1234123")) // [4]

	fmt.Println(UniqueNumber("76523752")) // [6 3]

	fmt.Println(UniqueNumber("12345")) // [1 2 3 4 5]

	fmt.Println(UniqueNumber("1122334455")) // []

	fmt.Println(UniqueNumber("0872504")) // [8 7 2 5 4]

}