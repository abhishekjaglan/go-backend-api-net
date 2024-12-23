package main

import "fmt"

type gasEngine struct {
	kmpl   int8
	litres int8
}

type electricEngine struct {
	kwh    int8
	kmpkwh int8
}

// generics in structs
type car[T gasEngine | electricEngine] struct {
	owner  string
	engine T
}

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))

	var flaotSlice = []float32{1, 2, 3, 4}
	fmt.Println(sumSlice(flaotSlice))

	var float64Slice = []float64{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(float64Slice))

	var myGasCar = car[gasEngine]{
		owner: "Abhishek",
		engine: gasEngine{
			kmpl:   20,
			litres: 10,
		},
	}

	var myElectricCar = car[electricEngine]{
		owner: "Abhishek",
		engine: electricEngine{
			kwh:    50,
			kmpkwh: 18,
		},
	}

	fmt.Println(myGasCar)
	fmt.Println(myElectricCar)
}

// generics in functions
func sumSlice[T int | float32 | float64](slice []T) T { // rather than making 3 different functions for 3 different data types we make one generic fn for all
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
