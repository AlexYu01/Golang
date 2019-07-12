/*
Returns the sign of the integer passed. + for positive numbers, - for negative numbers and 0 for the number 0.
*/
package main

func sign(n int) rune {
	if n > 0 {
		return '+'
	} else if n < 0 {
		return '-'
	} else {
		return '0'
	}
}
