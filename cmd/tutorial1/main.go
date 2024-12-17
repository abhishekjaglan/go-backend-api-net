package main

import (
	"fmt"
	"unicode/utf8"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello world")

	// arithmetic between diff var sizes
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// float numbers are ROUNDED DOWN when using int
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	//string
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	var str string = ` Hellioooo 
	worldiessss`
	fmt.Println(str)

	myString = " Hello \nWorld"
	fmt.Println(myString)

	fmt.Println("Length using inbuild len func in bytes", len(myString))

	fmt.Println("Actual length using utf8.RuneCountInString", utf8.RuneCountInString(myString))

	// Single quotes mean a rune
	var myRune rune = 'a'
	fmt.Println("My rune (Prints the ASCII value of 'a')", myRune)
	// fmt.Println("My rune len", utf8.RuneCount(myRune)) -> gives error

	// Inferred types
	var myStr = "New String"
	fmt.Println(myStr)

	NewStr := "Another New String" // Walrus Operator
	fmt.Println("Type Inferred using walrus operator", NewStr)

	var var1, var2 int = 1, 2
	fmt.Println(var1, var2)

	var returnType = sum(1, 2) // while this fine, its not evident whats the return typeof sum and therefore the value type in returnType. Hence, mention the return in variable for better reading of code
	fmt.Println(returnType)

	var returnType2 int = sum(2, 3)
	fmt.Println(returnType2)

	// constants
	const pii float32 = 3.14
	fmt.Println(pii)
	// const myConst int --> This gives error because cant leave const dcelared without initialization
}
