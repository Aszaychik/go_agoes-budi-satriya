package main

import (
	"fmt"
	"math"
)


func frog(jumps []int) int {
	cost := make([]int, len(jumps))

	// Inisialisasi biaya awal dari batu pertama ke kedua.
	cost[0] = 0
	cost[1] = int(math.Abs(float64(jumps[1] - jumps[0]))) // Menggunakan fungsi Abs untuk menghitung jarak antara 2 posisi

	// Iterasi untuk menghitung biaya minimum untuk mencapai setiap batu.
	for i := 2; i < len(jumps); i++ {
		// Biaya jika melompat dari batu sebelumnya.
		cost[i] = cost[i-1] + int(math.Abs(float64(jumps[i] - jumps[i-1])))
		
		// Biaya jika melompat dari batu sebelum batu sebelumnya.
		if cost[i-2] + int(math.Abs(float64(jumps[i] - jumps[i-2]))) < cost[i] {
			cost[i] = cost[i-2] + int(math.Abs(float64(jumps[i] - jumps[i-2])))
		}
	}

	return cost[len(jumps)-1]
}


func main() {

fmt.Println(frog([]int{10, 30, 40, 20})) // 30

fmt.Println(frog([]int{30, 10, 60, 10, 60, 50})) // 40

}