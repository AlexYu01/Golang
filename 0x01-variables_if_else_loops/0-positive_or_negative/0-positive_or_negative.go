/*
Generate pseudo random numbers and determine if they are posivite, negative, or zero.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var random int
	min := -100
	max := 100

	rand.Seed(time.Now().UnixNano())
	random = rand.Intn(max-min) + min
	fmt.Printf("The number %d is ", random)
	if random > 0 {
		fmt.Println("positive.")
	} else if random < 0 {
		fmt.Println("negative.")
	} else {
		fmt.Println("zero.")
	}
}
