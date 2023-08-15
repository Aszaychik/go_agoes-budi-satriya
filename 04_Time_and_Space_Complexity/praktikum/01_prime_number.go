package main

import "fmt"

func PrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	fmt.Println(PrimeNumber(1000000007)) // true

	fmt.Println(PrimeNumber(13)) // true

	fmt.Println(PrimeNumber(17)) // true

	fmt.Println(PrimeNumber(20)) // false

	fmt.Println(PrimeNumber(35)) // false

}