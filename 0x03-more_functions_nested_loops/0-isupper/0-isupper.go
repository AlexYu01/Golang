package main

// isUpper checks if the character passed is an uppercase character.
// It returns true if the rune is an uppercase character, otherwise false.
func isUpper(c rune) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	} else {
		return false
	}
}
