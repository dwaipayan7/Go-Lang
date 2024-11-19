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

func performPostRequest(){
	 todo := Todo{
		UserID: 22,
		Title: "Dwaipayan",
		Completed: true,
	}

	jsonData, err :=json.Marshal(todo)

	if err != nil {
		fmt.Println("Error Marshalling",err)
		return
	}

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	res, err :=	http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Getting error",err)
	}
	defer res.Body.Close()

	data,_ := ioutil.ReadAll(res.Body)
	fmt.Println("Data: ", string(data))


}

func main() {
	fmt.Println("Learning CRUD in GO")

	performGetRequest()

	performPostRequest()

}
