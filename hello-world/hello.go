package main

import "fmt"

func Hello() string {
	return "Hello World"
}

func main() {
	// Calling the Hello function
	fmt.Println(Hello())

	// Playing with rune literals
	fmt.Println("\u0009SomeInformationHere")
	fmt.Println('\u0009')
}
