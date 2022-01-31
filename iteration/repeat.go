package iteration

import "strings"

func Repeat(character string, num int, upper string) string {
	// I guess instead of setting up an empty string, we can use the explicit version here
	var repeated string

	if upper == "Up" {
		for i := 0; i < num; i++ {
			repeated += strings.ToUpper(character)
		}
		return repeated
	}
	// for number; if i < 5; add number?...
	for i := 0; i < num; i++ {
		repeated += character
	}
	return repeated
}
