package main

import "fmt"

func Pow(x, n int) int {
	result := 1

	if n > 0 {
		for i := 0; i < n; i++ {
			if n % 2 == 0 {
				result *= x * x
				n /= 2
			} else {
				result *= x
			}
		}
	}

	return result
}

func main() {
	fmt.Println(Pow(2, 3)) // 8
	
	fmt.Println(Pow(5, 3)) // 125
	
	fmt.Println(Pow(10, 2)) // 100
	
	fmt.Println(Pow(2, 5)) // 32
	
	fmt.Println(Pow(7, 3)) // 343
}