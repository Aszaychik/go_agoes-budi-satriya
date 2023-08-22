package main

import (
	"fmt"
	"sort"
)


type pair struct {

name string

count int

}


func MostAppearItem(items []string) []pair {

	var result []pair

	var countMapItem map[string]int = make(map[string]int)

	for _, item := range items {
		countMapItem[item]++
	}

	for name, count := range countMapItem {
		result = append(result, pair{name: name, count: count})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].count == result[j].count {
			return result[i].name < result[j].name
		}
		return result[i].count < result[j].count
	})


	return result
}


func main() {

	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()


	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()


	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // basketball->1 football->1 tenis->1

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()


}