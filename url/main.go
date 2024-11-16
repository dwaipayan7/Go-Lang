package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Dwaipayan URL")
	myUrl := "https://www.dwaipayan.com"
	fmt.Printf("%T",myUrl)

	parsedURL, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}

	fmt.Printf("%T",parsedURL)




}