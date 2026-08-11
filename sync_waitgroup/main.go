package main

import (
	"fmt"
	"sync"
)

// This file demonstrates the use of sync.WaitGroup for synchronizing goroutines
// WaitGroup is used to wait for a collection of goroutines to finish executing

// worker is a function that runs as a goroutine and represents concurrent work
func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	fmt.Printf("Worker %d starting\n", i)
	// some task is happening
	fmt.Printf("Worker %d done\n", i)
}

// main function demonstrates how to use WaitGroup to coordinate multiple goroutines
func main() {
	fmt.Println("Explore goroutine started")

	// Create a WaitGroup instance to track running goroutines
	var wg sync.WaitGroup
	// start multiple goroutines to run concurrently and wait for all to complete
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter before launching each goroutine
		go worker(i, &wg)
	}
	wg.Wait()                                 // Blocks until all goroutines have called Done() and counter reaches zero
	fmt.Println("Explore goroutine finished") // Prints after all goroutines complete
}
