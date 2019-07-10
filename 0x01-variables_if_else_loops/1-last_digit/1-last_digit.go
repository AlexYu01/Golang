/*
Generate pseudo random numbers and perform modulus in Go to get the last digit.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var random int
	min := 0
	max := 100

	rand.Seed(time.Now().UnixNano())
	random = rand.Intn(max-min) + min
	fmt.Printf("The last digit of %d is %d.\n", random, random % 10)
}