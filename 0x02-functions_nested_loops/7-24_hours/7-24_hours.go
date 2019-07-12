/*
Prints the 24 hour clock
*/
package main

import (
	"fmt"
)

func clock() {
	for hour := 0; hour < 24; hour++ {
		for min := 0; min < 60; min++ {
			fmt.Printf("%02d:%02d\n", hour, min)
		}
	}
}
