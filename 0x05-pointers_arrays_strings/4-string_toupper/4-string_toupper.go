package main

import (
	"strings"
)

// toUpper converts lower case alphabet letters in a string into upper case alphabet letters.
func toUpper(src1 string) string {
	var newStr strings.Builder

	for _, str := range src1 {
		if str >= 'a' && str <= 'z' {
			newStr.WriteRune(str - 32)
		} else {
			newStr.WriteRune(str)
		}
	}

	return newStr.String()
}
