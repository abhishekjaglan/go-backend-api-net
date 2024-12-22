package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed) // this output is because UTF-8 encoding is used by go, and french e is out of bound of ASCII characters, threfore UTF-8 uses more than 1 byte for string character 'french e' and it spills over for the capabilities of string
	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("The length of 'myString' is: %v \n", len(myString)) // returns no of bytes, therefor more than 6 as 2 e spill over to 2 extra bytes

	// to solve this issue, runes are used
	var myStringRune = []rune("résumé")
	var indexedRune = myStringRune[1]
	fmt.Printf("%v, %T\n", indexedRune, indexedRune)
	for i, v := range myStringRune {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of 'myStringRune' is %v: \n", len(myStringRune))
	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	var strSlice = []string{"j", "a", "g", "l", "a", "n"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i] // a new string is created everytime this operation takes place - inefficient and slow
	}
	fmt.Printf("%v\n", catStr)

	// sol to above - strings package
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("%v\n", catStr2)
}
