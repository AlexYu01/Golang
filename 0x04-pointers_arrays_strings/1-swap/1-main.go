package main

import (
	"fmt"
)

func main() {
	var x, y int = 10, 5

	fmt.Printf("x contains %d and y contains %d\n", x, y)
	fmt.Println("Swapping the values of x and y")
	swapInt(&x, &y)
	fmt.Printf("x contains %d and y contains %d\n", x, y)
}
