package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go is a statically typed, compiled programming language designed at Google. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.")
	fmt.Println("and it is often referred to as Golang due to its domain name, golang.org.")

	// myutil.PrintMessage("Hello, World!")

	var message string = "This is a message from the main package."
	var version = "1.0.0"
	fmt.Println(message)
	fmt.Println("Version:", version)

	var number int = 42
	fmt.Println("The answer to life, the universe, and everything is:", number)

	var dimension float64 = 3.14
	fmt.Println(dimension)

	var isActive bool = true
	fmt.Println("Is active:", isActive)

	var person = "John Doe"
	fmt.Println("Person:", person)

	const pi = 3.14159
	fmt.Println("The value of pi is:", pi)

	personAge := 30
	fmt.Println("Person's age:", personAge)

	var Public = "This is a public variable."
	var private = "This is a private variable."
	fmt.Println(Public)
	fmt.Println(private)
}
