package main

import (
	"fmt"
	"time"
)

// This file demonstrates how goroutines run concurrently with time.Sleep delays

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(2000 * time.Millisecond) // Simulate a delay
	fmt.Println("Hello again after 2 seconds!")
}

func sayHi() {
	fmt.Println("Hi, there!")
	time.Sleep(1000 * time.Millisecond) // Simulate a delay
	fmt.Println("Hi again after 1 second!")
}

func sayGoodbye() {
	fmt.Println("Goodbye for now!")
	time.Sleep(1500 * time.Millisecond) // Simulate a delay
	fmt.Println("Goodbye again after 1.5 seconds!")
}

func main() {
	fmt.Println("Learning goroutines in golang")
	go sayHello()
	go sayHi()
	go sayGoodbye()

	// wait a moment to allow the goroutines to finish executing before the main function exits
	time.Sleep(2200 * time.Millisecond) // Simulate a delay
}
