package main

import "fmt"

func main() {
	//array (fixed size in GO)

	// var intArr [3]int -> initialization
	var intArr2 [3]int = [3]int{1, 2, 3} // initialization and declaration
	intArr1 := [3]int{1, 2, 3}           // initialization and declaration
	intArr := [...]int{1, 2, 3, 4, 5}    // initilization and declaration
	//intArr[1] = 123
	fmt.Println(intArr[1])
	fmt.Println(intArr[0:3])
	fmt.Println(len(intArr))
	fmt.Println(intArr)
	fmt.Println(intArr1)
	fmt.Println(intArr2)

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// slices (wrapper on arrays - no fixed size)
	var intSlice []int = []int{4, 5, 6, 7}
	fmt.Println("intSlice", intSlice)
	intSlice = append(intSlice, 8)

	intSlice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("intSlice2", intSlice2)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice2), cap(intSlice2))
	intSlice2 = append(intSlice2, 10)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice2), cap(intSlice2))
	fmt.Println("intSlice2 after append", intSlice2)

	var intSlice3 []int = []int{8, 9}
	intSlice = append(intSlice, intSlice3...)
	fmt.Println(intSlice)

	// memory allocation for slice using make(type, len, cap)
	var strSlice []string = make([]string, 2, 4)
	strSlice = append(strSlice, "a", "b", "c", "d")
	fmt.Println(strSlice)
}
