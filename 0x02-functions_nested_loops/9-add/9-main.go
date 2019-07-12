package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num1 int
	var num2 int
	min := -100
	max := 100

	rand.Seed(time.Now().UnixNano())
	num1 = rand.Intn(max-min) + min
	num2 = rand.Intn(max-min) + min

	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, add(num1, num2))
}
