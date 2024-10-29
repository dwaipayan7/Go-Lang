package main

import "fmt"


func PublicFunction(){
	fmt.Println("This is public function")
}


func privateFunction(){
	fmt.Println("This is privateFunction")
}

func main() {
	// fmt.Println("Dwaipayan Biswas")

	// var name string = "Dwaipayan"
	// var version = "1.0"
	// var data = 90

	// fmt.Println(version)
	// fmt.Println(name)
	// fmt.Println(data)

	// var money int = 7000
	// fmt.Println(money)

	// var dimension float64 = 90.98
	// fmt.Println(dimension)

	// var decided bool = false
	// fmt.Println(decided)

	// var person = "Tatay Biswas"
	// fmt.Println(person)

	person := "123 dwaipayan"
	fmt.Println(person)

	var Public = "This is public"
	var private = "This is private"

	fmt.Println(Public)
	fmt.Println(private)

	PublicFunction()
	privateFunction()

}
