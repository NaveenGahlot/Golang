package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// This file demonstrates how to perform basic CRUD operations using HTTP methods in Go.
// CRUD stands for Create, Read, Update, and Delete, which are the four basic operations for managing resources in a web application.
// In this example, we will focus on the Read operation using an HTTP GET request to fetch data from a public API.

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func perforGetReuest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error in GET request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("GET request failed with status:", res.Status)
		return
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println("GET Response:", string(data))

	// Decode the JSON response into a Todo struct
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}
	fmt.Printf("Todo Item: %+v\n", todo)
	fmt.Println("Todo:-", todo)

	// Print individual fields of the Todo struct
	fmt.Println("UserID:-", todo.UserID)
	fmt.Println("ID:-", todo.ID)
	fmt.Println("Title:-", todo.Title)
	fmt.Println("Completed:-", todo.Completed)
}

func performPostRequest() {
	// Example of performing a POST request to create a new resource
	todo := Todo{
		UserID:    1,
		Title:     "New Todo Item",
		Completed: false,
	}
	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	// convert json data to string
	jsonString := string(jsonData)
	fmt.Println("JSON Data to be sent in POST request:", jsonString)

	// convert string data to reader
	reader := strings.NewReader(jsonString)
	myURL := "https://jsonplaceholder.typicode.com/todos"

	// Perform the POST request
	res, err := http.Post(myURL, "application/json", reader)
	if err != nil {
		fmt.Println("Error in POST request:", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response:-", string(data))
	fmt.Println("Response Status:-", res.Status)
}

func main() {
	fmt.Println("CRUD Operations Demo")
	// perforGetReuest()
	performPostRequest()

	// Note: The above code demonstrates the Read operation of CRUD. To implement Create, Update, and Delete operations, you would typically use HTTP POST, PUT/PATCH, and DELETE methods respectively, along with appropriate request bodies and endpoints.
}
