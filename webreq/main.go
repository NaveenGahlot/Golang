package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// This file demonstrates how to make an HTTP GET request in Go.
// Definition: This file contains the main function, which is the entry point of the program.

func main() {
	fmt.Println("learning web services")

	// Send an HTTP GET request to a sample API
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response", err)
		return
	}

	// Close the response body after use
	defer res.Body.Close()
	fmt.Printf("Type of response:- %T\n", res)
	// fmt.Println("Response:-", res)

	// Read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return
	}
	fmt.Println("response:-", string(data))
}
