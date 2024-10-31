package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("What is your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Print("Hello, Mr.", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ :=  reader.ReadString('\n')
	println("Hello ,Mr. ",name)


}
