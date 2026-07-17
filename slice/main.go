package main

import (
	"fmt"
)

func main() {
	// number := []int{1, 2, 3}
	// number = append(number, 6, 7, 8, 9)
	// fmt.Println("Number is:-", number)
	// fmt.Printf("Number has data type:- %T\n", number)
	// fmt.Println("length is:-", len(number))

	// name := []string{}
	// fmt.Println("name is:-", name)

	number := make([]int, 3, 5)
	number = append(number, 4)
	number = append(number, 98)
	number = append(number, 99)
	number = append(number, 4)
	number = append(number, 98)
	number = append(number, 99)
	number = append(number, 99)
	number = append(number, 99)
	number = append(number, 99)
	number = append(number, 99, 100, 101, 102, 103, 104)
	fmt.Println("Slice:-", number)
	fmt.Println("Length:-", len(number))
	fmt.Println("Capacity:", cap(number))
}
