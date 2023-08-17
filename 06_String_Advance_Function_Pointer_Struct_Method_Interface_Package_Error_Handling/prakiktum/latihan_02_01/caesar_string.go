package main

import (
	"fmt"
	"unicode"
)


func Caesar(offset int, input string) string {
	var result string

	for _, char := range input {
		if unicode.IsDigit(char) { // Cek jika karakter adalah digit/number
			result += string('0' + (char -'0'+ int32(offset))%10) // Menambah offset karakter digit antara 0 - 9 (%10)
		} else if char >= 'a' && char <= 'z' { // Cek jika karakter adalah huruf kecil
			result += string('a' + (char -'a'+ int32(offset))%26) // Menambah offset karakter lower case antara a - z (%26)
		} else if char >= 'A' && char <= 'Z' { // Cek jika karakter adalah huruf besar
			result += string('A' + (char -'A'+ int32(offset))%26) // Menambah offset karakter upper case antara A - Z (%26)
		} else { // Jika karakter bukan huruf/digit
			result += string(char)
		}
	}

	return result
}


func main() {

fmt.Println(Caesar(3, "abc")) // def

fmt.Println(Caesar(2, "alta")) // cnvc

fmt.Println(Caesar(10, "alterraacademy")) // kvdobbkkmknowi

fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

fmt.Println(Caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}