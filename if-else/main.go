package main

import "fmt"

// if-else condition in go lang
func main() {
	x := 2
	if x > 5 { // if conditon
		fmt.Println("x is greater than 5")
	} else { // else conditon
		fmt.Println("x is less than or equal to 5")
	}

	// nested if-else if-else condition
	z := 10
	if z > 10 { // if condition
		fmt.Println("z is greater than 5")
	} else if z > 5 { // else if condition
		fmt.Println("z is greater than 5 but smaller than 10")
	} else { // else condition
		fmt.Println("z is smaller then 5")
	}

	y := 10
	if x < 5 && (y > 5 || z > 5) {
		fmt.Println("hey how are you")
	} else {
		fmt.Println("let's go to the else block")
	}
}
