package main

import (
	"fmt"
	"os"
	"time"
)

func printXMultiples(x int) error  {
	if x <= 0 {
		return fmt.Errorf("Invalid input: number must be greater than 0")
	}

	for multiple := 1; ; multiple++ {
		fmt.Printf("Multiple of %d times %d: %d\n", x, multiple, x * multiple)
		time.Sleep(3 *time.Second)
	}
}

func main() {
	var input int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	go func() {
		err = printXMultiples(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}
	}()

	time.Sleep(15 * time.Second)
	fmt.Println("Done")
}