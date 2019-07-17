package main

import (
	"fmt"
)

// printSkip prints one char out of 2 of a string, followed by a new line.
func printSkip(aString *string) {
	for i := 0; i < len(*aString); i += 2 {
		fmt.Printf("%c", (*aString)[i])
	}
	fmt.Println()
}
