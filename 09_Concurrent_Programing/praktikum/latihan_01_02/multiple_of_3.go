package main

import (
	"fmt"
)

func printMultipleOf3(numberOfTimes int) error {
	if numberOfTimes <= 0 {
		return fmt.Errorf("Invalid input : Input must be greater than 0")
	}

	for i := 1; i <= numberOfTimes; i++ {
		fmt.Printf("Multiple of 3 Times %d: %d\n", i, i*3)
	}

	return nil
}

func main() {
	var numberOfTimes int
	fmt.Print("Enter the number of times to print multiple of 3: ")
	_, err := fmt.Scan(&numberOfTimes)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	errChan := make(chan error)
	
	go func() {
		defer close(errChan)
		err := printMultipleOf3(numberOfTimes)
		errChan <- err
	}()

	err = <-errChan
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}


	fmt.Println("Done")
}