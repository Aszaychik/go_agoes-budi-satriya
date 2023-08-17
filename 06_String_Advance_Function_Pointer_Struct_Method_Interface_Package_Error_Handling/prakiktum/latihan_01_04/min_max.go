package main

import (
	"fmt"
	"sort"
)


func GetMinMax(numbers ...*int) (min, max int) {
	// Membuat array baru dengan panjang sama dengan jumlah parameter
	var arrayNumbers []int = make([]int, len(numbers))

	// Memasukkan data parameter ke dalam array
	for i, number := range numbers {
		arrayNumbers[i] = *number
	}

	// Menggunakan sort library untuk mengurutkan array
	sort.Ints(arrayNumbers)

	// Mengambil nilai minimal (index pertama) dan maksimal (index terakhir)
	min = arrayNumbers[0]
	max = arrayNumbers[len(arrayNumbers)-1]

	return min, max
}


func main() {

var a1, a2, a3, a4, a5, a6, min, max int

fmt.Scan(&a1)

fmt.Scan(&a2)

fmt.Scan(&a3)

fmt.Scan(&a4)

fmt.Scan(&a5)

fmt.Scan(&a6)

fmt.Println()


min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

fmt.Println("Nilai Max : ", max)

fmt.Println("Nilai Min : ", min)

}

