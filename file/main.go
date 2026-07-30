package main

import (
	"fmt"
	"io"
	"os"
)

// This file demonstrates how to create a new file in Go, write content into it, and handle errors during file operations.
// It is a simple example of file handling using os.Create, io.WriteString, and defer to close the file safely.
// Important points: create the file, write text, check for errors, and print success or failure messages.

func main() {
	// This program creates a file named example.txt and writes a message into it.
	// It also shows how to handle errors while creating and writing the file.

	// Create a new file for writing.
	file, err := os.Create("example.txt")
	if err != nil {
		// Print an error message if the file cannot be created.
		fmt.Println("Error while creating file:-", err)
		return
	}

	// Close the file automatically when the function ends.
	defer file.Close()

	// Define the text content that will be written into the file.
	content := "Hello world"

	// Write the content to the file and get the number of bytes written.
	byte, errors := io.WriteString(file, content+"\n")
	fmt.Println("Byte return it:-", byte)

	// Check whether the write operation was successful.
	if errors != nil {
		fmt.Println("Error while writing file:-", errors)
		return
	}

	// Print a success message after the file is written.
	fmt.Println("Successfully created file")

	/*
		// Example: open the created file and read it line by line.
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while opening file:-", err)
			return
		}
		defer file.Close()

		// Create a buffer to read the file content.
		buffer := make([]byte, 1024)

		// Read the file content into the buffer.
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file:-", err)
				return
			}

			// Print the content that was read from the file.
			fmt.Println(string(buffer[:n]))
		}
	*/

	/*
		// Example: read the whole file into a byte slice in one step.
		content, err := ioutil.ReadFile("example.txt")
		if err != nil {
			fmt.Println("Error while reading file:-", err)
			return
		}
		fmt.Println(string(content))
	*/
}
