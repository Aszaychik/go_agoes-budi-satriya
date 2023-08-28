package main

import "fmt"

func toBinary(number int) []string {
	var result []string = make([]string, number+1) // membuat slice dengan panjang number + 1

	for i := 0; i <= number; i++ {
		result[i] = fmt.Sprintf("%b", i) // penggunaan format %b untuk konversi ke biner 
	}

	return result
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	var toBinaryResult = toBinary(number)

	fmt.Printf("The Binary of %d is %v\n", number, toBinaryResult)
}