package main

import (
	"strings"
)

// strNCat concatenates two strings where n number of runes are copied from source. The new string will be returned as a result.
func strNCat(src1, src2 *string, n int) string {
	var newStr strings.Builder
	var strLen int = len(*src2)

	for _, str := range *src1 {
		newStr.WriteRune(str)
	}

	for i := 0; i < n && i < strLen; i++ {
		newStr.WriteRune(rune((*src2)[i]))
	}

	return newStr.String()
}
