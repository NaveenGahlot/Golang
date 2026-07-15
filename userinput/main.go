package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, What's your name?")
	// var name string
	// fmt.Scanln(&name)
	// fmt.Println("Hey", name, "How old are you?")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hey", name, "How old are you?")
	age, _ := reader.ReadString('\n')
	fmt.Println("Hey", name, "You are", age, "years old")
}
