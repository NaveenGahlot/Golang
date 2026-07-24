package main

import "fmt"

// Pointer is a variable that stores the memory address of another variable,
// rather than storing the data directly. Use the asterisk (*) to declare and
// dereference a pointer, and the ampersand (&) to get the address of a value.

// modifyValueByReference multiplies the integer value stored at the provided
// pointer by 20.
func modifyValueByReference(num *int) {
	*num = *num * 20
}

func main() {
	// Create a pointer to an integer value.
	var num int = 2
	var ptr *int = &num

	// Print the integer value, the pointer address, and the dereferenced pointer value.
	fmt.Println("Number has value:-", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value at address ptr:", *ptr)

	// A nil pointer is a valid pointer value that points to nothing.
	var ptr2 *int
	if ptr2 == nil {
		fmt.Println("Pointer is not assigned")
	}

	// Pass the variable by reference so the function can modify the original value.
	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value contains:-", value)
}
