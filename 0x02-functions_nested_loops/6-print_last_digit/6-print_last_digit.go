/*
Returns the last digit in the integer
*/
package main

func lastDigit(n int) int {
	if n < 0 {
		n *= -1
	}
	return n % 10
}
