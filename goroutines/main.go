package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Function ended")
}

func tatay() {
	fmt.Println("Tatay")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("Dwaipayan")
	go sayHello()
	go tatay()

	time.Sleep(800 * time.Millisecond)
}
