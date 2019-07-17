package main

import (
	"fmt"
)

func printHalf(aStr *string) {
	fmt.Println((*aStr)[len(*aStr)/2:])
}
