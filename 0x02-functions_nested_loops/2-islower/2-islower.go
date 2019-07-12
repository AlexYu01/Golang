/*
Determines if the passed rune variable is a lower case alphabet letter.
*/
package main

import (
	"fmt"
)

func isLower(c rune) {
	if c < 'a' || c > 'z' {
		fmt.Println("not a lower case alphabet.")
	} else {
		fmt.Println("a lower case alphabet.")
	}
}
