package main

import "fmt"

func main() {
	var myString string = "Hello World"
	printMe(myString)

	var numerator int = 11
	var denominator int = 2
	var result, remainder int = intDivision(numerator, denominator)
	fmt.Println("Result = ", result)
	fmt.Println("Remainder = ", remainder)
	fmt.Printf("intDivision(%v,%v) gives Result = %v, and Remainder = %v\n", numerator, denominator, result, remainder)
}

func printMe(myString string) {
	fmt.Println(myString)
}
func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
