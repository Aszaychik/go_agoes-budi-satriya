package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var mergedArrayInput []string = append(arrayA, arrayB...) // merge arrays with unknown size

	var uniqueValues map[string]bool = make(map[string]bool) // create map for unique Values
	var result []string = []string{}

	for _, value := range mergedArrayInput {
		if !uniqueValues[value] {
			uniqueValues[value] = true
			result = append(result, value)
		}
	}

	return result
}

func main() {

	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []

}