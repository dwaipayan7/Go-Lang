package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 4, 66, 3)
	// fmt.Println("Number ", numbers)
	// fmt.Printf("Number has data type of: %T\n", numbers)


	numbers := make([]int,3,5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 5)
	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 5)
	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 5)

	fmt.Println("Slice ",numbers)
	fmt.Println("Length", len(numbers))
	fmt.Println("Capacity", cap(numbers))

}
