package main

import "fmt"


func Fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	return Fibonacci(number - 1) + Fibonacci(number - 2)
}

func main() {

	fmt.Println(Fibonacci(0)) // 0

	fmt.Println(Fibonacci(2)) // 1

	fmt.Println(Fibonacci(9)) // 34

	fmt.Println(Fibonacci(10)) // 55

	fmt.Println(Fibonacci(12)) // 144

}