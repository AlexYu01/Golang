package main

import (
	"fmt"
)

func printTenTimes() {
	for i := 0; i < 10; i++ {
		for c := 'a'; c < 'z'; c++ {
			fmt.Printf("%c ", c)
		}
		fmt.Println("z")
	}
}
