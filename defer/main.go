package main

import "fmt"

// The defer keyword schedules a function call to run after the current function completes.
// It is useful for cleanup tasks, closing files, or printing final messages.
// It uses a stack-like structure, so the last deferred function runs first.

// add takes two integers and returns their sum.
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of the program")

	// Call the add function and store the result in a variable.
	result := add(5, 6)

	// defer delays the execution of these statements until main finishes.
	defer fmt.Println("Data is:", result)
	defer fmt.Println("Middle of the program")

	fmt.Println("End of the program")
}
