package main

import (
	"fmt"
)

func Compare(a, b string) string {
	var result string

	// Mencari panjang string terkecil
	var minStringLength int = len(a)
	if len(b) < minStringLength {
		minStringLength = len(b)
	}

	// Perulangan untuk mendapatkan value index string
	for i := 0; i < minStringLength; i++ {
		// jika value index string sama, tambahkan ke result
		if a[i] == b[i] {
			result += string(a[i])
		} else {
			// break jika value index string tidak sama
			break
		}
	}

	return result
}

func main() {

fmt.Println(Compare("AKA", "AKASHI")) //AKA

fmt.Println(Compare("KANGOORO", "KANG")) //KANG

fmt.Println(Compare("KI", "KIJANG")) //KI

fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU

fmt.Println(Compare("ILALANG", "ILA")) //ILA

}

