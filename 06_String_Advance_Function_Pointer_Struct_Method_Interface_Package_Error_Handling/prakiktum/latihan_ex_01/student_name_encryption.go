package main

import (
	"fmt"
)


type Student struct {

	name string

	nameEncode string

	score int

}


type Chiper interface {

	Encode() string

	Decode() string

}

// Table subtitusi sesuai contoh di pada latihan explorasi, sisanya ditambahkan secara acak
var substitutionDict map[byte]byte = map[byte]byte{
	'r': 'i', 'i': 'r', 'z': 'a', 'k': 'p', 'y': 'b','a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v', 'f': 'u',
	'g': 't', 'h': 's', 'j': 'q', 'l': 'o', 'm': 'n', 'n': 'm',
	'o': 'l', 'q': 'j', 's': 'h', 't': 'g', 'u': 'f', 'v': 'e',
	'w': 'd', 'x': 'c',
	'R': 'I', 'I': 'R', 'Z': 'A', 'K': 'P', 'Y': 'B','C': 'X', 'D': 'W', 'E': 'V', 'F': 'U',
	'G': 'T', 'H': 'S', 'J': 'Q', 'L': 'O', 'M': 'N', 'N': 'M',
	'O': 'L', 'Q': 'J', 'S': 'H', 'T': 'G', 'U': 'F', 'V': 'E',
	'W': 'D', 'X': 'C',
}

func (s *Student) Encode() string {
	var encodedName string
	for i := 0; i < len(s.name); i++ { 
		var char byte = s.name[i]
		encodedChar, exists := substitutionDict[char] // mencari key karakter apakah karakter ada di key map
		if exists { // jika ada, maka value dari key tersebut akan ditambahkan ke variabel 
			encodedName += string(encodedChar) 
		} else { // jika tidak ada, maka karakter akan ditambahkan ke variabel (seperti : angka (123) atau special character (#@) )
			encodedName += string(char)
		}
	}

	return encodedName
}

func (s *Student) Decode() string {
	var decodedName string = ""
	for i := 0; i < len(s.name); i++ { 
		var char byte = s.name[i]
		var found bool = false // sebagai flag, jika true maka decodeName tidak akan ditambahkan ke variabel
		for k, v := range substitutionDict { // looping map untuk mencari key yang sama dengan karakter yang diinput
			if char == v { // jika key yang sama dengan karakter yang diinput ada di map maka value dari key tersebut akan ditambahkan ke variabel
				decodedName += string(k)
				found = true // set variable found menjadi true, agar decode name yg ketemu tidak di tambahkan lagi
				break 
			}
		}
		if !found { // jika found false maka karakter akan ditambahkan ke variabel (seperti : angka (123) atau special character (#@) )
			decodedName += string(char)
		}
	}

	return decodedName
}


func main() {

	var menu int

	var a Student = Student{}

	var c Chiper = &a


	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")

	fmt.Scan(&menu)


	if menu == 1 {

		fmt.Print("\nInput Student Name: ")

		fmt.Scan(&a.name)

		fmt.Print("\nEncode of student's name " + a.name + " is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student Name: ")

		fmt.Scan(&a.name)

		fmt.Print("\nDecode of student's name " + a.name + " is : " + c.Decode())

	}

}