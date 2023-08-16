package main

import "fmt"

func GradeCalculation(nilai int) string {
	switch {
	case nilai >= 80 && nilai <= 100:
		return "A"
	case nilai >= 65 && nilai < 80:
		return "B"
	case nilai >= 50 && nilai < 65:
		return "C"
	case nilai >= 35 && nilai < 50:
		return "D"
	case nilai >= 0 && nilai < 35:
		return "E"
	default:
		return "Nilai Invalid"
	}
}

func main() {
	var nilai int
	fmt.Println("Masukkan nilai: ")
	fmt.Scan(&nilai)

	var result string = GradeCalculation(nilai)

	fmt.Println(result)
}
