package main

import (
	"encoding/json"
	"fmt"
)

// Person defines a simple data structure used to demonstrate JSON encoding and decoding.
type Person struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	IsAdult string `json:"is_adult"`
}

// main shows how to convert a struct to JSON and back to a struct.
func main() {
	fmt.Println("learning json in golang")
	persion := Person{
		Name:    "John Doe",
		Age:     "30",
		IsAdult: "true",
	}
	fmt.Println("persion data:-", persion)

	// convert person struct to JSON Encoding (Marshal)
	jsonData, err := json.Marshal(persion)
	if err != nil {
		fmt.Println("error in json encoding:-", err)
		return
	}
	fmt.Println("json data:-", string(jsonData))

	// convert JSON data to person struct (Unmarshal)
	var person2 Person
	err = json.Unmarshal(jsonData, &person2)
	if err != nil {
		fmt.Println("error in json decoding:-", err)
		return
	}
	fmt.Println("person2 data:-", person2)
}
