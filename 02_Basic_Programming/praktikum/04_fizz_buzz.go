package main

import "fmt"

func FizzBuzz(n int) []string {
	results := make([]string, n)
	for i := 1; i <= n; i++ {
			if i%3 == 0 && i%5 == 0 {
					results[i-1] = "FizzBuzz"
			} else if i%3 == 0 {
					results[i-1] = "Fizz"
			} else if i%5 == 0 {
					results[i-1] = "Buzz"
			} else {
					results[i-1] = fmt.Sprintf("%d", i)
			}
	}
	return results
}


func main() {
	result := FizzBuzz(100)

	fmt.Println(result)
}