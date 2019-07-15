package main

import (
	"fmt"
)

func main() {
	var n int = 49
	var c int = 120

	fmt.Printf("%d is ", n)
	if isDigit(n) {
		fmt.Println("a digit in ascii.")
	} else {
		fmt.Println("not a digit in ascii.")
	}
	fmt.Printf("%d is ", c)
	if isDigit(c) {
		fmt.Println("a digit in ascii.")
	} else {
		fmt.Println("not a digit in ascii.")
	}
}
