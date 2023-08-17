package main

import "fmt"

type Car struct {
	carType string
	fuelin  float32
}

func (car *Car) EstimatedDistanceTraveled() string {
	var result string

	const fuelConsumption float32 = 1.5 // konsumsi bahan bakar 1.5 L / mill

	var estDistance float32 = car.fuelin * fuelConsumption // Hitung perkiraan jarak dengan mengalikan jumlah bahan bakar dengan konsumsi bahan bakar

	result = fmt.Sprintf("Car type: %s, Est. distance: %.2f", car.carType, estDistance) // menggunakan Sprintf untuk mengembalikan nilai string, %s untuk menyimpan string, %.2f untuk menyimpan float dengan hanya menampilkan 2 angka desimal 

	return result
}

func main() {
	var car Car = Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	var result string = car.EstimatedDistanceTraveled()

	fmt.Println(result)
}