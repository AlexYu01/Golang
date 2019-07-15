package main

import (
	"fmt"
)

// printNum prints a number without the usage of Printf
func printNum(n int) {
	var m int

	if n < 0 {
		fmt.Print("-")
		m = -n
	} else {
		m = n
	}
	printRecur(m)
}

// printRecur prints a number using recursion
func printRecur(m int) {
	if m/10 != 0 {
		printRecur(m / 10)
	}
	fmt.Print(rune(m % 10))
}
