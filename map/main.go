package main

import "fmt"

// A map is a collection of unordered key-value pairs. In Go, maps are implemented as hash tables. The keys in a map are unique, and each key maps to exactly one value. Maps are useful for situations where you need to associate data with unique identifiers, such as storing student grades by their names.
func main() {
	// name <-> grade
	studentGrades := make(map[string]int) // Declare a map with string keys and int values

	// Adding key-value pairs to the map
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 92
	studentGrades["David"] = 88
	studentGrades["Eve"] = 95

	fmt.Println("Marks of Bob:-", studentGrades["Bob"])         // Accessing the value associated with the key "Bob"
	studentGrades["Bob"] = 95                                   // Updating the value associated with the key "Bob"
	fmt.Println("Updated Marks of Bob:-", studentGrades["Bob"]) // Accessing the updated value associated with the key "Bob"

	// Deleting a key-value pair from the map
	delete(studentGrades, "Bob")
	fmt.Println("Marks of Bob after deletion:-", studentGrades["Bob"])

	// Check if a key exists in the map
	grade, exists := studentGrades["David"]
	fmt.Println("Grade of David:", grade)
	fmt.Println("Does David exist in the map?", exists)

	// fmt.Println("Marks of Bob:-", studentGrades["Bob"])

	// Iterating over the map using a for loop
	fmt.Println("All student grades:")
	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	// Another way to declare and initialize a map
	fmt.Println("\nAnother way to declare and initialize a map:-")
	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 92,
		"David":   88,
		"Eve":     95,
	}
	for index, value := range person {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}
}
