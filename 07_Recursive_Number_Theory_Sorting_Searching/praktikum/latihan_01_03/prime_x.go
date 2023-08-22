package main

import (
	"fmt"
	"math"
)


func isPrime(number int) bool {
	if number <= 1 {
			return false
	}
	sqrt := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrt; i++ {
			if number % i == 0 {
					return false
			}
	}
	return true
}

func PrimeX(number int) int {
	var count, i int = 0, 2
	
	for {
		if isPrime(i) {
			count++
		}

		if count == number {
			return i
		}

		i++
	}
}


func main() {

	fmt.Println(PrimeX(1)) // 2

	fmt.Println(PrimeX(5)) // 11

	fmt.Println(PrimeX(8)) // 19

	fmt.Println(PrimeX(9)) // 23

	fmt.Println(PrimeX(10)) // 29

}

