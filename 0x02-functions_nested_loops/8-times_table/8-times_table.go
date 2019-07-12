/*
Prints the multiplication table starting from 0 to 9
*/
package main

import (
	"fmt"
)

func timesTable() {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			fmt.Printf("%2d", row*col)
			if col != 9 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}
