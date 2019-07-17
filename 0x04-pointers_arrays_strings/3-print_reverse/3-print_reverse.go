package main

import (
	"fmt"
)

// printStr takes a string pointer and prints the string rune by rune in reverse
func printRev(str *string) {

	for i := len(*str) - 1; i >= 0; i-- {
		fmt.Printf("%c", (*str)[i])
	}
	fmt.Println()
}
