/*
Use a for-loop to print the alphabet in lower case.
*/
package main

import (
	"fmt"
)

func main() {
	for i := 'a'; i < 'z'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println("z")
}