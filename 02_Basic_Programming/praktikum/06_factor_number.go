package main

import "fmt"

func FactorNumber(number int) []int {
	factors := []int{}
	for i := 1; i <= number; i++ {
			if number % i == 0 {
					factors = append(factors, i)
			}
	}
	return factors
}


func main() {
	var number int
	fmt.Print("Input a number: ")
	fmt.Scan(&number)

	result := FactorNumber(number)
	fmt.Println(result)
}