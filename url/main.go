package main

import (
	"fmt"
	"net/url"
)

// This file demonstrates how to parse and work with URLs in Go.
// A URL (Uniform Resource Locator) is the web address used to locate a resource.

// This section shows how to parse a URL and read its parts.
func main() {
	fmt.Println("Learning URL")

	// Example URL string
	myURL := "https://jsonplaceholder.typicode.com/todos/1"
	// fmt.Printf("Type of URL:- %T\n", myURL)

	// Parse the URL into a structured object
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}
	// fmt.Printf("Type of URL:- %T\n", parsedURL)

	// Print important parts of the URL
	fmt.Println("Scheme:-", parsedURL.Scheme)
	fmt.Println("Host:-", parsedURL.Host)
	fmt.Println("Path:-", parsedURL.Path)
	fmt.Println("RawQuery:-", parsedURL.RawQuery)

	// Modify URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=mkjaat"

	// Construct a URL string from the updated URL object
	newUrl := parsedURL.String()
	fmt.Println("new URl:-", newUrl)
}
