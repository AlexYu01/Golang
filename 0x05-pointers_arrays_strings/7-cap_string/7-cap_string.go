package main

import (
	"strings"
)

// capString takes in a string and capitalizes all words of a string. There is
// also a function called `Title` in the strings package that does this.
// Returns the resulting string.
func capString(src string) string {
	var new strings.Builder
	var prev rune

	prev = ' '
	for _, char := range src {
		if char >= 'a' && char <= 'z' {
			if sep(prev) {
				new.WriteRune(char - 32)
			} else {
				new.WriteRune(char)
			}
		} else {
			new.WriteRune(char)
		}
		prev = char
	}

	return new.String()
}

// sep takes in a rune and determines if it matches a separator rune.
// Returns True if the rune is a separator, False otherwise.
func sep(char rune) bool {
	separator := []rune(" \t\n,;.!?\"(){}")

	for _, s := range separator {
		if char == s {
			return true
		}
	}

	return false
}
