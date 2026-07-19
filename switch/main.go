package main

import "fmt"

// This program demonstrates the use of switch statements in Go. It checks the value of a variable and executes different code blocks based on that value.
func main() {
	day := 3
	switch day { // switch statement to check the value of day
	case 1: // case for Monday
		fmt.Println("Monday") // print Monday if day is 1
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default: // default case if day is not 1, 2, or 3
		fmt.Println("Another day") // print Another day if day is not 1, 2, or 3
	}

	// Switch statement with multiple cases for months
	month := "January"
	switch month {
	case "January", "February", "March":
		fmt.Println("It's the first quarter of the year.")
	case "April", "May", "June":
		fmt.Println("It's the second quarter of the year.")
	case "July", "August", "September":
		fmt.Println("It's the third quarter of the year.")
	case "October", "November", "December":
		fmt.Println("It's the fourth quarter of the year.")
	default:
		fmt.Println("Invalid month.")
	}

	// Switch statement with a condition.
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("It's freezing!")
	case temperature >= 0 && temperature < 20:
		fmt.Println("It's cold.")
	case temperature >= 20 && temperature < 30:
		fmt.Println("It's warm.")
	default:
		fmt.Println("It's hot!")
	}
}
