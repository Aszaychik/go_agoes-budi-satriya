package main

import (
	"fmt"
	"strings"
)

// Table konversi number ke romawi
type RomanTable []struct {
	Value int // Nilai desimal
	Digit string // Simbol Romawi
}

func toRomawi(number int) string {
	const maxRomanNumber int = 3999 // Nilai maksimum angka yang dapat dikonversi ke Romawi
	var result string

	if number > maxRomanNumber {
		result = "Exceeds the maximum Roman number"
		return result
	}

	var numberToRomanTable RomanTable = RomanTable{ // Menggunakan Table konversi number ke romawi
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder // Membuat variabel yang akan menampung hasil konversi number ke romawi dan menggunakan strings.buillder
	for _, conversion := range numberToRomanTable {
		// Melakukan konversi angka desimal ke simbol Romawi
		for number >= conversion.Value {
			roman.WriteString(conversion.Digit)
			number -= conversion.Value
		}
	}

	return roman.String()
}

func main() {
	var number int
	fmt.Print("Input a number : ")
	fmt.Scan(&number)
	
	var convertToRoman string = toRomawi(number)

	fmt.Println(convertToRoman) // Output: "XXIII"
}
