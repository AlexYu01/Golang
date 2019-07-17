package main

import (
	"fmt"
)

func main() {
	num := 50
	var p *int

	p = &num

	fmt.Printf("This int pointer points to the memory space %p containing the value %d\n", p, *p)
	resetTo98(p)
	fmt.Printf("Changing the value in the memory space at %p\n", p)
	fmt.Printf("This int pointer points to the memory space %p containing the value %d\n", p, *p)
}
