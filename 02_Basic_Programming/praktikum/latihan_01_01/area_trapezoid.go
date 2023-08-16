package main

import "fmt"

func AreaOfTrapezoid(base1 float32, base2 float32, height float32) float32 {
	return 0.5 * (base1 + base2) * height
}

func main() {
	fmt.Println("Masukkan base1: ")
	var base1 float32
	fmt.Scan(&base1)

	fmt.Println("Masukkan base2: ")
	var base2 float32
	fmt.Scan(&base2)

	fmt.Println("Masukkan height: ")
	var height float32
	fmt.Scan(&height)

	var result float32 = AreaOfTrapezoid(base1, base2, height)

	fmt.Println(result)
}