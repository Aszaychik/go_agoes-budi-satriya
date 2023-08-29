package main

import "fmt"


func fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	var a int = 0
	var b int = 1
	
	for i := 2; i <= number; i++ {
		var c = a + b
		a = b
		b = c
	}

	return b
}

func main() {

	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144

}