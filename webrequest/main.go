package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Go Language")

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("Getting error")
		return
	}
	defer res.Body.Close()

	fmt.Printf("Type of response: %T\n", res)
	//Read Response

	data,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}

	fmt.Println("Response: ",string(data))

}
