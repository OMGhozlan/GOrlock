package toiterate

import "strings"

// Repeat returns a string with the character repeated count times.
func Repeat(character string, count int) (repeated string) {
	var i int
	for i = 0; i < count; i++ {
		repeated += character
	}
	return
}

func Repeat2(char string, count int) (repeated string) {
	repeated = strings.Repeat(char, count)
	return
}
