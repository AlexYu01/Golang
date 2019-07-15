package main

// isDigit Checks if the integer passed is an ascii value that corresponds to the characters 0 through 9.
// It returns true if the integer represents a digit in ascii, otherwise false.
func isDigit(n int) bool {
	if n >= '0' && n <= '9' {
		return true
	} else {
		return false
	}
}