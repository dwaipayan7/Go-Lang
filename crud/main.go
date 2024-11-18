package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

func main() {
	fmt.Println("Learning CRUD in GO")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
	fmt.Printf("Getting Error\n")
	return
	}
	defer res.Body.Close()

	data,err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Getting Error\n")
		return
	}

	fmt.Println("Data: ",string(data))


}