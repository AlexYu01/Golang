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

	fmt.Printf("%d has a %c sign\n", random, sign(random))
}
