/*
Returns the absolute value of a the integer passed
*/
package main

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}
