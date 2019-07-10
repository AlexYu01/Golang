/*
Prints all numbers from 00 to 99 in ascending order
*/
package main

import (
	"fmt"
)

func main() {
	j := 0
	for i := 0; i < 10; i++ {
		for ; j < 10; j++ {
			fmt.Printf("%d%d", i, j)
			if j == 9 {
				fmt.Println()
			} else {
				fmt.Printf(", ")
			}
		}
		j = 0
	}
}
