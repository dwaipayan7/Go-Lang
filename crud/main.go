package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {

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

	fmt.Println("Completed response: ", todo.Completed)

}

func performPostRequest() {
	todo := Todo{
		UserID:    22,
		Title:     "Dwaipayan",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error Marshalling", err)
		return
	}

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	res, err := http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Getting error", err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Data: ", string(data))

}

func performUpdateRequest() {
	todo := Todo{
		UserID:    8,
		Title:     "Dwaipayan Biswas",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error Marshalling", err)
		return
	}

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	//create put request
	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT Request", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))
	fmt.Println("Response status: ", res.Status)

}

func performDeleteRequest() {

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)

	if err != nil {
		fmt.Println("Error creating DELETE Request", err)
		return
	}

	// req.Header.Set("Content-type", "application/json") --> Not required

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status: ", res.Status)

}

func main() {
	fmt.Println("Learning CRUD in GO")

	// performGetRequest()

	// performPostRequest()

	// performUpdateRequest()

	performDeleteRequest()

}
