package main // package declaration

import "fmt" // import statement to include the fmt package for formatted I/O

// This is a simple program to the use of functions in Golang
func simpleFunction() { // function definition
	fmt.Println("This is a simple function") // function body
}

// This function takes two integers as parameters and returns their sum
func add(a int, b int) int { // function definition with parameters and return type
	return a + b // function body with return statement
}

func subtract(a, b int) int { // function definition with parameters and return type
	result := a - b // function body with assignment to local variable
	return result   // return statement
}

// This function takes two integers as parameters and returns their product
func multiply(a, b int) (result int) { // function definition with named return value
	result = a * b // function body with assignment to named return value
	return         // return statement (optional for named return values)
}

// The main function is the entry point of the program
func main() {
	fmt.Println("We are learning functions in Golang") // print statement
	simpleFunction()                                   // function call

	result := add(5, 10)               // function call with return value
	fmt.Println("The sum is:", result) // print statement with return value

	asn0 := subtract(15, 5)                 // function call with return value
	fmt.Println("The difference is:", asn0) // print statement with return value

	asn1 := multiply(20, 30)             // function call with named return value
	fmt.Println("The product is:", asn1) // print statement with named return value
}
