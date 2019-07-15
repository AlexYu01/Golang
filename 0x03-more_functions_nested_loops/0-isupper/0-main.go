package main

import (
	"fmt"
)

func main() {
	var c rune = 'A'
	var n rune = '0'

	fmt.Printf("%c is ", c)
	if isUpper(c) {
		fmt.Println("an upper case character.")
	} else {
		fmt.Println("not an upper case character.")
	}
	fmt.Printf("%c is ", n)
	if isUpper(n) {
		fmt.Println("an upper case character.")
	} else {
		fmt.Println("not an upper case character.")
	}
}
