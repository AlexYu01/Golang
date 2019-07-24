package main

import (
	"fmt"
)

func main() {
	src := "ROT13 (\"rotate by 13 places\", sometimes hyphenated ROT-13) is a simple letter substitution cipher.\n"

	fmt.Print(src)
	fmt.Print(rot13(src))
}
