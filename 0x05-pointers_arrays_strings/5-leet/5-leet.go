package main

import (
	"strings"
)

// leet encodes a string into 1337 code.
func leet(src string) string {
	var lower, upper, number = "aeotl", "AEOTL", "43071"
	var newStr strings.Builder
	var j int

	for _, str := range src {
		for j = 0; j < 5; j++ {
			if str == rune(lower[j]) || str == rune(upper[j]) {
				newStr.WriteRune(rune(number[j]))
				break
			}
		}
		if j == 5 {
			newStr.WriteRune(str)
		}
	}

	return newStr.String()
}
