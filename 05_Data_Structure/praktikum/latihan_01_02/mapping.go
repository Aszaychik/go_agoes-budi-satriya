package main

import "fmt"

func Mapping(slice []string) map[string]int {
	var indexCount map[string]int = make(map[string]int)

	for _, value := range slice {
		indexCount[value]++
	}

	return indexCount
}

func main() {

	// Test cases

	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"})) // map[asd:2 qwe:1]

	fmt.Println(Mapping([]string{})) // map[]

}