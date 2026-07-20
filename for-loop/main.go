package main

import "fmt"

// This is a simple program to demonstrate the use of for loop in Golang
func main() {
	for i := 0; i < 1; i++ { // for loop with initialization, condition and increment
		fmt.Println("Numbers is:-", i)
	}

	// for loop with only condition
	counter := 0
	for {
		fmt.Println("Counter is:-", counter)
		counter++
		if counter == 1 {
			break
		}
	}

	// for loop with range
	numbers := []int{1, 2, 3, 4, 5}     // slice of integers
	for index, value := range numbers { // for loop with range to iterate over slice
		fmt.Printf("Index is:- %d and Value is:- %d\n", index, value)
	}

	// for loop with range to iterate over string
	data := "Hello World"            // string data
	for index, value := range data { // for loop with range to iterate over string
		fmt.Printf("Index of Data is:- %d and value is:- %c\n", index, value)
	}
}
