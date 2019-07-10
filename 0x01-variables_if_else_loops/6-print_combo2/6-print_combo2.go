/*
Prints all two digit combinations except for the inverse of the lowest combination. eg. 01 will be printed but not 10. 00, 11, 22 etc. are also not printed.
*/
package main

import (
	"fmt"
)

func main() {
	j := 1
	for i := 0; i < 10; i++ {
		for ; j < 10; j++ {
			fmt.Printf("%d%d", i, j)
			if j == 9 {
				fmt.Println()
			} else {
				fmt.Printf(", ")
			}
		}
		j = i + 2
	}
}
