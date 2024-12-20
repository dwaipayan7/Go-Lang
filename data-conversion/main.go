package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Number is: ", num)
	fmt.Printf("Type of num is: %T\n", num)
	// num = num + 56.7

	var data float64 = float64(num)
	data = data + 56.7
	fmt.Println("Data is: ", data)
	fmt.Printf("Type of Data is: %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is: ", str)
	fmt.Printf("Type of str is: %T\n", str)

	number_String := "12345"
	number_Integer, _ := strconv.Atoi(number_String)

	fmt.Println("Number of Integer is: ", number_Integer)
	fmt.Printf("Type of Number of Integer is: %T\n", number_Integer)

	num_string := "3.14"
	number_float,_ := strconv.ParseFloat(num_string,64)
	fmt.Println("Number of Float is: ",number_float)
	fmt.Printf("Type of number_float is: %T\n",	number_float)

}
