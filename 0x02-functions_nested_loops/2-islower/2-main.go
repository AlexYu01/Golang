package main

import (
	"fmt"
)

func main() {

	var first, second, third rune = 'a', 'N', '1'
	fmt.Printf("%c is ", first)
	isLower(first)
	fmt.Printf("%c is ", second)
	isLower(second)
	fmt.Printf("%c is ", third)
	isLower(third)
}
