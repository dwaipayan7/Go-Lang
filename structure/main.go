package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// fmt.Println("Dwaipayan")

	var dwaipayan Person
	fmt.Println("Dwaipayan Person: ",dwaipayan)

	dwaipayan.FirstName = "Dwaipayan"
	dwaipayan.LastName = "Biswas"
	dwaipayan.Age = 22
	fmt.Println("Dwaipayan Person: ",dwaipayan)


	person1 := Person{
		FirstName: "Tatay",
		LastName: "Biswas",
		Age: 22,
	}

	fmt.Println("Person 1: ",person1)

}
