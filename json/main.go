package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	fmt.Println("Dwaipayan Json Data")

	person := Person{Name: "Dwaipayan", Age: 22, IsAdult: true}
	fmt.Println(person)

	//convert person into Json Encode (Marshalling)
	jsonData,err := json.Marshal(person);

	if err != nil {
		fmt.Printf("Error in Marshalling: %v\n", err)
	}
	fmt.Println("Person Data is:",string(jsonData))

	//Decoding
	var personData Person
	err = json.Unmarshal(jsonData, &personData)

}
