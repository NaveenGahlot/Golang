package main

import (
	"fmt"
	"time"
)

// Package main demonstrates basic time operations in Go.

func main() {
	// Get the current date and time
	currentTime := time.Now()
	fmt.Println("Current time:-", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)

	// Format the current time in a readable custom layout
	formatted := currentTime.Format("02-01-2006, Monday, 15:04:05, 3:04 PM")
	fmt.Println("Formatted time:-", formatted)

	// Define a layout and a date string, then parse the string into a time value
	layoutStr := "2006-01-02"
	dateStr := "2025-11-25"
	formattedTime, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Formatted Time:-", formattedTime)

	// Add one day to the current time
	newDate := currentTime.Add(24 * time.Hour)
	fmt.Println("New date time:-", newDate)

	// Format the new date with a different layout
	formattedNewDate := newDate.Format("2006/01/02, Monday")
	fmt.Println("Formatted new date:-", formattedNewDate)
}
