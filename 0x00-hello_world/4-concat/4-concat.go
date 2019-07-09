/*
Prints a concatenated string using fmt's Println func
*/
package main

import "fmt"

func main() {
	var first string = "I am the first string."
	var second string = "Well I am the second string."
	var chained string = first + " " + second
	fmt.Println(chained)
}
