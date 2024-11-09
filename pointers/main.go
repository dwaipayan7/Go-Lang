package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 20
} 

func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := 2
	ptr := &num

	// fmt.Println("Number Value is: ", num)
	fmt.Println("Value is: ",ptr)
	fmt.Println("Data Contains: ",*ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value is: ",value)


}

