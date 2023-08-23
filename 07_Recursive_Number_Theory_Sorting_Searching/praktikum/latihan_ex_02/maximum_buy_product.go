package main

import (
	"fmt"
	"sort"
)


func MaximumBuyProduct(money int, productPrice []int) int {
	var result int

	sort.Slice(productPrice, func(i, j int) bool { // Sorting productPrice dari yg terkecil
		return productPrice[i] < productPrice[j]
	})

	for _, price := range productPrice {
		if money >= price { // Jika money lebih besar dari price maka bisa membeli barang
			money -= price
			result ++ // Sebuah counter untuk menghitung jumlah barang yang dibeli
		} else {
			break
		}

	}

	return result
}


func main() {

	fmt.Println(MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})) // 3

	fmt.Println(MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})) // 4

	fmt.Println(MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})) // 4

	fmt.Println(MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})) // 1

	fmt.Println(MaximumBuyProduct(0, []int{10000, 30000})) // 0

}