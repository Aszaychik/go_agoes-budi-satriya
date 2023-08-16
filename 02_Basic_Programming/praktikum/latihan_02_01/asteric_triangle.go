package main

import "fmt"

func AstericTriangle(rows int) string {
	var result string

	if rows <= 0 {
		result = "Rows must be greater than 0"
		return result
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			result += " "
		}
		for j := 1; j <= i; j++ {
			result += "* "
		}
		result += "\n"
	}

	return result
}

func main() {
	var result string = AstericTriangle(5)

	fmt.Println(result)
}