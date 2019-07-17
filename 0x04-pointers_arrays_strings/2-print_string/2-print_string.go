package main

import (
	"fmt"
)

// printStr takes a string pointer and prints the string rune by rune
func printStr(str *string) {
	strLen := len(*str)

	for i := 0; i < strLen; i++ {
		fmt.Printf("%c", (*str)[i])
	}
	fmt.Println()
}
