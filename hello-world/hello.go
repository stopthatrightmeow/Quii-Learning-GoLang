package main

import (
	"fmt"
)

// Constants, variables that are.. constant...
const spanish = "Spanish"
const french = "French"
const german = "German"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	// Setting default prefix, and over riding that with a switch case.
	// Something that's not in Python but will be neat to learn none-the-less...
	prefix := englishHelloPrefix
	// Neat, switches...
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	}

	return prefix + name
}

func main() {
	// Calling the Hello function
	fmt.Println(Hello("", ""))

}
