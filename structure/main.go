package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// fmt.Println("Dwaipayan")

	//1st method
	var dwaipayan Person
	fmt.Println("Dwaipayan Person: ", dwaipayan)

	dwaipayan.FirstName = "Dwaipayan"
	dwaipayan.LastName = "Biswas"
	dwaipayan.Age = 22
	fmt.Println("Dwaipayan Person: ", dwaipayan)

	//2nd method
	person1 := Person{
		FirstName: "Tatay",
		LastName:  "Biswas",
		Age:       22,
	}

	type Contact struct {
		Email string
		Phone string
	}

	type Address struct {
		House int
		Area  string
		State string
	}

	type Employee struct {
		Person_Details Person
		Person_Contact Contact
		Person_Address Address
	}

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Dwaipayan",
		LastName:  "Biswas",
		Age:       22,
	}

	employee1.Person_Contact.Email = "123@gmail.com"
	employee1.Person_Contact.Phone = "7550"

	employee1.Person_Address.Area = "Barasat"
	employee1.Person_Address.House = 123
	employee1.Person_Address.State = "West Bengal"

	fmt.Println("Employee 1: ->", employee1)

	fmt.Println("Person 1: ", person1)

	//3rd method

	var person2 = new(Person)
	person2.FirstName = "Hello"
	person2.LastName = "World"
	person2.Age = 22

	fmt.Println("Person 2: ", person2)

}
