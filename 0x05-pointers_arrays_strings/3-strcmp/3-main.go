package main

import (
	"fmt"
)

func main() {
	var s1, s2 = "Hello", "World!"

	fmt.Printf("%d\n", strCmp(s1, s2))
	fmt.Printf("%d\n", strCmp(s2, s1))
	fmt.Printf("%d\n", strCmp(s1, s1))
}
