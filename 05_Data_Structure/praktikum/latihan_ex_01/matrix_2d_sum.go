package main

import (
	"fmt"
	"math"
)

func Matrix2dSum(matrix [][]int) int {
    // Buat variabel `difference` untuk menyimpan hasil penjumlahan
    var difference int = 0

    // Iterasi setiap baris dalam matrix
    for index, row := range matrix {
        // Hitung perbedaan antara elemen pertama dan terakhir pada baris
        difference += row[index] - row[len(matrix)-index-1]
    }

    // Konversi `difference` ke tipe integer dan simpan ke variabel `absoluteDifference`
    var absoluteDifference int = int(math.Abs(float64(difference)))

    // Kembalikan `absoluteDifference`
    return absoluteDifference
}

func main() {
    var matrix [][]int = [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{9, 8, 9},
    }

    var result int = Matrix2dSum(matrix)

    fmt.Println("Absolute Difference: ", result)
}