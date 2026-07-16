package main

import "fmt"

// error handling in the functions
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("division by zero is not allowed")
// 	}
// 	return a / b, nil
// }

// error handling in the functions but returning string instead of error
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "division by zero is not allowed"
	}
	return a / b, "nil"
}

func main() {
	fmt.Println("Started Error Handling in the functions")
	// err handling in the functions but returning error
	// answer, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("The answer is:", answer)

	// err handling in the functions but use underscore instead of error
	answer, _ := divide(10, 4)
	fmt.Println("The answer is:", answer)
}
