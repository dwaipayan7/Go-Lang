package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting the program")
	data := add(12, 8)
	defer fmt.Println("Data is: ", data)
	defer fmt.Println("Middle of the program")
	fmt.Println("Last of the program")
}
