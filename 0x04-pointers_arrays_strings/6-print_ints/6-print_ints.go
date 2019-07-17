package main

import (
	"fmt"
)

// printArray takes a slice to an array of integers and prints n elements
func printArray(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d", arr[i])
		if i < n-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
