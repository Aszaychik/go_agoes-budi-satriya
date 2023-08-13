package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsPalindrome(words string) bool {
	words = strings.ToLower(words)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		if words[i] != words[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Apakah Palindrome?")
	fmt.Print("Masukkan kata: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if IsPalindrome(input) {
		fmt.Printf("Input: %s\nPalindrome\n", input)
	} else {
		fmt.Printf("Input: %s\nBukan Palindrome\n", input)
	}
}
