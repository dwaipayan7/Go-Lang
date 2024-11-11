package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time: ", currentTime)
	fmt.Printf("Type of currentTime: %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006, 3:04 PM ,Monday")
	fmt.Println("Formatted time: ", formatted)

	layout_str := "2006-01-02"
	date_str := "2023-11-25"
	formatter_time, _ := time.Parse(layout_str, date_str)
	fmt.Println("Formatted Time: ", formatter_time)

	//add 1 day to the current time

	new_date := currentTime.Add(72 * time.Hour)
	fmt.Println("New Date time: ", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("Formatted New Date time: ", formatted_new_date)

}
