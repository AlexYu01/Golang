/*
Print the alphabets except for 'm' and 'u'
*/
package main

import (
	"fmt"
)

func main() {
	for i := 'a'; i < 'z'; i++ {
		if i != 'm' && i != 'u' {
			fmt.Printf("%c ", i)
		}
	}
	fmt.Println("z")
}