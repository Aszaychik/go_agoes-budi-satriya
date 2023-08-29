package main

import "fmt"

var fibonacciCache = map[int]int{}


func fibonacci(number int) int {
	if number <= 1 {
		fibonacciCache[number] = number
		return number
	}

	if index, found := fibonacciCache[number]; found {
		return index
	}

	fibonacciCache[number] = fibonacci(number-1) + fibonacci(number-2)

	return fibonacci(number-1) + fibonacci(number-2)
}

func main() {

	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144

}