package main

import (
	"fmt"
	"strconv"
)

/*
This file demonstrates basic data conversion in Go.
It shows how to convert values between different data types
using type conversion and the strconv package.
*/

func main() {
	// Integer value declaration and printing.
	var num int = 20
	fmt.Println("Number is:-", num)
	fmt.Printf("Type of num is %T\n", num)

	// Convert an integer to a float64 value.
	var data float64 = float64(num)
	fmt.Println("data is:-", data)
	fmt.Printf("Type of Data is %T\n", data)

	// Convert an integer to a string using strconv.Itoa.
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is:-", str)
	fmt.Printf("Type of str is %T\n", str)

	// Convert a string to an integer using strconv.Atoi.
	number_string := "123"
	number_int, _ := strconv.Atoi(number_string)
	number_int = number_int + 456
	fmt.Println("number_int is:-", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	// Convert a string to a float64 using strconv.ParseFloat.
	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("number_float is:-", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)
}
