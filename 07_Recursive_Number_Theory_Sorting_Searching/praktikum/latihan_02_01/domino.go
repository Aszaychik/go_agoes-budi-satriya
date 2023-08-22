package main

import "fmt"


func PlayingDomino(cards [][]int, deck []int) interface{} {
	for _, card := range cards {
		matchesCardWithDeckHead := card[0] == deck[0] || card[1] == deck[0]
		matchesCardWithDeckTail := card[0] == deck[1] || card[1] == deck[1]

		if matchesCardWithDeckHead || matchesCardWithDeckTail {
			return card
		}
	}
	
	return "Tutup kartu"
}


func main() {

	fmt.Println(PlayingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})) // [3, 4]

	fmt.Println(PlayingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})) // [6, 5]

	fmt.Println(PlayingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1})) // "Tutup kartu"

}