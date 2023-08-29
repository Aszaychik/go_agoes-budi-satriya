package main

import "fmt"


func simpleEquations(a, b, c int) bool {
	var foundSolution bool = false

	for x := 1; x <= 10000; x++ {
		for y := 1; y <= 10000; y++ {
			var z int = a - x - y
			if z <= 0 {
				continue
			}

			if x*y*z == b && x*x+y*y+z*z == c {
				fmt.Println(x, y, z)
				foundSolution = true
				return foundSolution
			}
		}
	}

	if !foundSolution {
		fmt.Println("no solution")
	}
	return foundSolution
}



func main() {

	simpleEquations(1, 2, 3) // no solution

	simpleEquations(6, 6, 14) // 1 2 3

}

