package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Tatay"] = 100
	studentGrades["Dwaipayan"] = 99
	studentGrades["Hello"] = 90

	fmt.Println("Marks of Tatay is: ", studentGrades["Tatay"])

	studentGrades["Tatay"] = 99
	fmt.Println("Updated Marks: ", studentGrades["Tatay"])

	delete(studentGrades, "Hello")

	grades, exist := studentGrades["Hello"]

	fmt.Println("Grade of Hello", grades)
	fmt.Println("Does Hello exist in the map?", exist)

	fmt.Println(studentGrades["Hello"])

	for index, value := range studentGrades {

		fmt.Println("Student Name: ", index, " Marks: ", value)
	}
	println()

	person := map[string]int{
		"ALice":     90,
		"Tatay":     80,
		"Dwaipayan": 88,
	}

	for index, value := range person {
		fmt.Println("Person Name: ", index, " Marks:", value)
	}

}
