package main

import (
	"fmt"
)

// printNums prints numbers from 0 to 9.
func printNums() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}
