package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}

	return "Odd"
}

func main() {
	var number int
	fmt.Println("Input number : ")
	fmt.Scan(&number)

	result := EvenOrOdd(number)

	fmt.Println(result)
}