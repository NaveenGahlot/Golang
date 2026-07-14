package main

import "fmt"

func main() {
	age := 25
	name := "John Doe"
	height := 5.65556

	fmt.Println("age:-", age, "name:-", name, "height:-", height)
	fmt.Println("hello world")

	fmt.Printf("age is:- %d\n name:- %s\n height:- %.4f\n", age, name, height)
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of height is %T\n", height)
}
