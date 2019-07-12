/*
Determines if the passed rune variable is a alphabet letter.
*/
package main

import (
	"fmt"
)

func isAlpha(c rune) {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		fmt.Println("an alphabet letter.")
	} else {
		fmt.Println("not an alphabet letter.")
	}
}
