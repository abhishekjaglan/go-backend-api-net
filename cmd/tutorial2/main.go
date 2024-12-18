package main

import (
	"errors"
	"fmt"
)

func main() {
	var myString string = "Hello World"
	printMe(myString)

	var numerator int = 11
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator) // error handling as well

	// error handling and if else statements
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Println("Remainder = ", remainder)
		fmt.Println("Result = ", result)
	} else {
		fmt.Printf("intDivision(%v,%v) gives Result = %v, and Remainder = %v\n", numerator, denominator, result, remainder)
	}

	// switch
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Println("Remainder = ", remainder)
		fmt.Println("Result = ", result)
	default:
		fmt.Printf("intDivision(%v,%v) gives Result = %v, and Remainder = %v\n", numerator, denominator, result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("Remainder 0")
	case 1, 2:
		fmt.Println("Remainder 1 or 2")
	}
}

func printMe(myString string) {
	fmt.Println(myString)
}
func intDivision(numerator int, denominator int) (int, int, error) { //error return declarartion
	// customn error generation
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by 0")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}
