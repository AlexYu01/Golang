package main

import (
	"strings"
)

// strCat concatenates two strings and returns the new string
func strCat(src1, src2 *string) string {
	var newStr strings.Builder

	for _, str := range *src1 {
		newStr.WriteRune(str)
	}

	for _, str := range *src2 {
		newStr.WriteRune(str)
	}
	
	return newStr.String()
}
