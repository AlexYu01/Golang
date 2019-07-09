/*
Practice slicing in Golang
*/
package main

import "fmt"

func main() {
	var sliceMe string = "I am a string to be sliced"
	fmt.Printf("Slice from index 0 to index 3(exclusive): %s\n", sliceMe[0:3])
	fmt.Printf("Slice from index 5 to index 15(inclusive): %s\n", sliceMe[5:16])
}