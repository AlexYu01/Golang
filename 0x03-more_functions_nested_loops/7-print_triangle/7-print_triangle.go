package main

import (
	"fmt"
)

// printTri prints a triangle using '#' in the terminal. Size of the triangle is determined by the integer passed.
func printTri(size int) {
	if size > 0 {
		for row := 0; row < size; row++ {
			for col := 0; col < size; col++ {
				if col >= (size - row - 1) {
					fmt.Print("#")
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
