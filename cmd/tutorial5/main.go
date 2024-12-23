package main

import "fmt"

type gasEngine struct {
	kpl    uint8
	litres uint8
	owner  ownerInfo
	yearBought
}

type yearBought struct {
	year string
}

type ownerInfo struct {
	name string
}

type electricEngine struct {
	kpkwh uint8
	kwh   uint8
}

func (e electricEngine) kilometersLeft() uint8 {
	return e.kpkwh * e.kwh
}

type engine interface {
	kilometersLeft() uint8
}

// receiver methods - tied to structs
func (e gasEngine) kilometersLeft() uint8 {
	return e.kpl * e.litres
}

func canMakeIt(e engine, kilometers uint8) { // interface is used to generalise use of kilometersLeft() from both structs at one place
	if e.kilometersLeft() >= kilometers {
		fmt.Println("You can make it there")
	} else {
		fmt.Println("Sorry, not enough gas")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{kpl: 25, litres: 5}
	fmt.Println(myEngine.kpl, myEngine.litres)
	myEngine.owner.name = "Abhishek"
	myEngine.year = "1989"
	fmt.Println(myEngine)
	myEngine.yearBought.year = "1898"
	fmt.Println(myEngine.owner.name)
	fmt.Println(myEngine.owner)
	fmt.Println(myEngine)
	fmt.Println(myEngine.kilometersLeft())

	// anonymous struct - initializing a struct without creating a specific type for it i.e on the fly
	myEngine2 := struct {
		litres uint
		kpl    uint
	}{25, 5}
	fmt.Println(myEngine2)

	canMakeIt(myEngine, 124)
	var myElectricEngine electricEngine = electricEngine{25, 5}
	canMakeIt(myElectricEngine, 150)
}
