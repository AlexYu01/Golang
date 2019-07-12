package main

import (
	"fmt"
)

func main() {

	var first, second, third rune = 'a', 'N', '1'
	fmt.Printf("%c is ", first)
	isAlpha(first)
	fmt.Printf("%c is ", second)
	isAlpha(second)
	fmt.Printf("%c is ", third)
	isAlpha(third)
}
