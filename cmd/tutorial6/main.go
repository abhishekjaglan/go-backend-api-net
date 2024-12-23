package main

import "fmt"

func main() {
	var p *int8 // assigns nil value to p as it not pointing anywhere
	//fmt.Println(*p) // gives runtime error :invalid memory address or nil pointer dereference

	p = new(int8)   // assigns a free memory location to p and 0 value at that location bcz no initialization
	fmt.Println(*p) // prints value of value stored in memory location in p and hence pointing to it
	fmt.Println(p)  //  prints memory address stored in p, just like any variable

	var ptr *int8
	var value int8 = 10
	ptr = &value // ptr stores memory address of value and now points to it
	fmt.Println(*ptr)
	*ptr = 34
	fmt.Println(*ptr)

	var thing1 [5]int8 = [5]int8{1, 2, 3, 4, 5}
	// calling function by value
	fmt.Printf("Memory location of thing1 array: %p\n", &thing1)
	var result [5]int8 = square(thing1)
	fmt.Println("Value of thing1 after calling function square: ", thing1)
	fmt.Println("Value of result array: ", result)
	// calling function by reference
	fmt.Println(&thing1[0])
	fmt.Println(&thing1[1])
	fmt.Printf("Memory location of thing1 array: %p\n", &thing1)
	result = squarePointer(&thing1)
	fmt.Println("Value of thing1 after calling function square Pointer: ", thing1)
	fmt.Println("Value of result array: ", result)
}

func square(thing2 [5]int8) [5]int8 { // additional space is required by thing2
	fmt.Printf("Memory location of thing2 %p\n", &thing2)
	for i := range thing2 {
		thing2[i] *= thing2[i]
	}
	return thing2
}

func squarePointer(thing2 *[5]int8) [5]int8 {
	fmt.Printf("Memory location of thing2 %p\n", &thing2)
	fmt.Println(&thing2[0])
	fmt.Println(*thing2)
	fmt.Println(&thing2[1])
	for i := range thing2 {
		thing2[i] *= thing2[i]
	}
	return *thing2
}
