package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)


func calculateFrequencyLetters(word string) map[string]int {
	result := make(map[string]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	isLetter := regexp.MustCompile(`^[A-Za-z]+$`).MatchString // Regexp untuk mengecek apakah string adalah huruf atau tidak


	for _, letter := range word {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			if isLetter(l) { // Jika huruf maka akan ditambahkan ke result
				mutex.Lock()
				result[l]++
				mutex.Unlock()
			}
		}(string(letter))
	}

	wg.Wait()
	return result
}

func main() {
	newScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the word: ")
	newScanner.Scan()
	inputWord := newScanner.Text()

	frequencyLetters := calculateFrequencyLetters(inputWord)

	for key, value := range frequencyLetters {
		fmt.Printf("%s: %d\n", key, value)
	}
}
