package main

import (
	"fmt"
	"strings"
)

// This file demonstrates basic string operations using Go's built-in strings package.
// It shows how to split, count, trim, and join strings in a simple executable program.

func main() {
	// Split a comma-separated string into a slice of substrings.
	data := "apple, orange, banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	// Count how many times the substring "two" appears in the string.
	str := "One two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("count:-", count)

	// Remove leading and trailing whitespace from the string.
	str = "  Hello         world  "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed:-", trimmed)

	// Join multiple strings together with a space separator.
	str1 := "Prince"
	str2 := "Agarwal"
	result := strings.Join([]string{str1, "Kumar", str2}, " ")
	fmt.Println("result:-", result)
}
