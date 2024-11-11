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

}
