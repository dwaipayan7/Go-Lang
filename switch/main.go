package main

import "fmt"

func main() {
	fmt.Println("Enter your number: ")
	var day int
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Wrong Input")
	}

	month := "January"

	switch month {
	case "January":
		fmt.Println("Winter")
	case "February":
		fmt.Println("Spring")

	default:
		fmt.Println("Wrong Input")
	}

	temperature := 25
	switch {
	case temperature <= 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

}
