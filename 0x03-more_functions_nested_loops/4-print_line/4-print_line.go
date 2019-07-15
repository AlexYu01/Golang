package main

import (
	"fmt"
)

// printLine prints a line in the terminal where length is determined by the integer passed.
func printLine(len int) {
	if len > 0 {
		for i := 0; i < len; i++ {
			fmt.Print("_")
		}
	}
	fmt.Println()
}
