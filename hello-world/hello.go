package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	// Calling the Hello function
	fmt.Println(Hello(""))

}
