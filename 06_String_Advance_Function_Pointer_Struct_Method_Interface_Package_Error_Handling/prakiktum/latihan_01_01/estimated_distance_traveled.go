package main

import "fmt"

type Car struct {
	carType string
	fuelin  float32
}

func (car *Car) EstimatedDistanceTraveled() float32 {
	var result float32

	const fuelConsumption float32 = 1.5 // konsumsi bahan bakar 1.5 L / mill

	result = car.fuelin * fuelConsumption

	return result
}

func main() {
	var car Car = Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	var estDistance float32 = car.EstimatedDistanceTraveled()

	fmt.Printf("Car type: %s, Est. distance: %.2f", car.carType, estDistance)// menggunakan Sprintf untuk mengembalikan nilai string, %s untuk menyimpan string, %.2f untuk menyimpan float dengan hanya menampilkan 2 angka desimal
}