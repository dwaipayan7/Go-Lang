package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD in GO")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Printf("Getting Error\n")
		return
	}
	defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Printf("Getting Error\n")
	// 	return
	// }

	// fmt.Println("Data: ", string(data))

	var todo Todo
	json.NewDecoder(res.Body).Decode(&todo)

	fmt.Println("Todo: ", todo)

	fmt.Println("Completed response: ",todo.Completed);

}
