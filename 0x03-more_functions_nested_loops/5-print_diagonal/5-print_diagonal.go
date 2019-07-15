package main

import (
	"fmt"
)

// printDiag prints a diagonal line in the terminal where length is determined by the integer passed.
func printDiag(len int) {
	if len > 0 {
		for row := 0; row < len; row++ {
			for col := 0; col < len; col++ {
				if row == col {
					fmt.Print("\\")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}
