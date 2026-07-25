package main

import (
	"fmt"
)

/*
This is the main entry point of the Go learning project.
It demonstrates basic Go syntax, variables, constants,
and different data types using simple print statements.
*/

func main() {
	// Introduction to the Go programming language.
	fmt.Println("Go is a statically typed, compiled programming language designed at Google. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.")
	fmt.Println("and it is often referred to as Golang due to its domain name, golang.org.")

	// myutil.PrintMessage("Hello, World!")

	// String and version variable examples.
	var message string = "This is a message from the main package."
	var version = "1.0.0"
	fmt.Println(message)
	fmt.Println("Version:", version)

	// Integer variable example.
	var number int = 42
	fmt.Println("The answer to life, the universe, and everything is:", number)

	// Float variable example.
	var dimension float64 = 3.14
	fmt.Println(dimension)

	// Boolean variable example.
	var isActive bool = true
	fmt.Println("Is active:", isActive)

	// String variable example.
	var person = "John Doe"
	fmt.Println("Person:", person)

	// Constant example.
	const pi = 3.14159
	fmt.Println("The value of pi is:", pi)

	// Short declaration example.
	personAge := 30
	fmt.Println("Person's age:", personAge)

	// Public and private variable example.
	var Public = "This is a public variable."
	var private = "This is a private variable."
	fmt.Println(Public)
	fmt.Println(private)
}
