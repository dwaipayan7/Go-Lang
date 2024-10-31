package main

import "fmt"

func simpleFunction() {
	fmt.Println("Dwaipayan")
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func multiply(a , b int) (result int){
	result = a * b
	return
}

func main() {
	fmt.Println("Functions in Go")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println(ans)

	mul :=  multiply(5, 6)
	fmt.Println(mul)

}
