/*
Prints all three digit combinations, but only the lowest value of the permutations (012 is printed but not 120, 102, 021 etc.) in ascending order.
*/
package main

import (
	"fmt"
)

func main() {
	j := 1
	k := 2
	for i := 0; i < 10; i++ {
		for ; j < 10; j++ {
			for ; k < 10; k++ {
				fmt.Printf("%d%d%d", i, j, k)
				if k == 9 {
					fmt.Println()
				} else {
					fmt.Printf(", ")
				}
			}
			k = j + 2
		}
		j = i + 1
	}
}
