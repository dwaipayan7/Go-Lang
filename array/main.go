package main

import (
	"fmt"
)

func main() {
	// fmt.Println("This is Array")

	// var name [5]string
	// name[0] = "Dwaipayan"
	// name[1] = "Biswas"

	// fmt.Println("Name is:", name)

	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(numbers)
	// fmt.Println(len(numbers))

	// fmt.Println("value of name at 2nd index is: ", name[0])

	var name [5]string

	name[0] = "Dwaipayan"
	name[1] = "Tatay"

	fmt.Println("Price is ", name)

	fmt.Printf("%q\n", name)
	fmt.Println(len(name))
}
