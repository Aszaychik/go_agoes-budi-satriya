package main

import "fmt"

func pascalTriangle(numRows int) [][]int {
	var result [][]int

	for i := 0; i < numRows; i++ {
		var row []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i { // Cek awal index dan akhir index
				row = append(row, 1) // Jika iya, tambahkan 1
			} else {
				row = append(row, result[i-1][j-1] + result[i-1][j]) // Jika tidak, tambahkan dari hasil penjumlahan dari index sebelumnya [i-1][j-1] dan [i-1][j]
			}
		}
		result = append(result, row)
	}

	return result
}


func main() {
	var numRows int
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&numRows)
	
	var pascalTriangleResult [][]int = pascalTriangle(numRows)

	fmt.Printf("Pascal Triangle of %d rows: %v\n", numRows, pascalTriangleResult)
}
