package main

// package array

import "fmt"

func main() {
	fmt.Println("array in golang")

	var name [5]string // declaring an array of strings with a size of 5
	name[0] = "John"   // assigning values to the array
	name[1] = "Doe"
	name[2] = "Jane"
	name[3] = "Smith"
	name[4] = "Alice"

	fmt.Println("The names in the array are:", name) // printing the array

	// declaring an array of integers with a size of 5
	var age = [5]int{1, 2, 3, 4, 5} // assigning values to the array
	// age[0] = 25
	// age[1] = 30
	// age[2] = 35
	// age[3] = 40
	// age[4] = 45

	fmt.Println("The ages in the array are:", age)           // printing the array
	fmt.Println("The length of the age array is:", len(age)) // printing the length of the array
	fmt.Println("value of name at  2 index is:", name[2])    // printing the value at index 2 of the array

	var names [5]string // declaring an array of string with a size of 5
	names[2] = "John"   // assigning value to the array at index 2
	names[0] = "Doe"    // assigning value to the array at index 0

	fmt.Println("names is:", names)          // printing the array
	fmt.Printf("names Array is %q\n", names) // printing the array with formatting
}
