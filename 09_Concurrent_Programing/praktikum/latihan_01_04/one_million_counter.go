package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup

func oneMilCounter() int {
	var result int

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				defer wg.Done()
				wg.Add(1)
				mutex.Lock()
				result ++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	return result
}

func oneMilCounterWithoutMutex() int {
	var result int

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				defer wg.Done()
				wg.Add(1)
				result ++
			}
		}()
	}

	wg.Wait()
	return result
}

func main() {
	counterWithoutMutex := oneMilCounterWithoutMutex()
	counter := oneMilCounter()
	fmt.Println("Counter with mutex : ", counter) // Counter :  1000000
	fmt.Println("Counter without mutex : ", counterWithoutMutex) // Counter :  RANDOM
}