/*
A function that prints the alphabet.
*/
package main

import (
	"fmt"
)

func printAlphabet() {
	for i := 'a'; i < 'z'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println("z")
}
