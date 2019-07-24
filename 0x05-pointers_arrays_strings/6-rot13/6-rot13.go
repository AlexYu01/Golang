package main

import (
	"strings"
)

// rot13 takes a string and encodes it using rot13. Returns the encoded rot13 string
func rot13(src string) string {
	var new strings.Builder
	var rotIdx int

	rotI := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rotO := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

	for _, char := range src {
		for rotIdx = 0; rotIdx < len(rotI); rotIdx++ {
			if char == rune(rotI[rotIdx]) {
				new.WriteRune(rune(rotO[rotIdx]))
				break
			}
		}
		if rotIdx == len(rotI) {
			new.WriteRune(char)
		}
	}

	return new.String()
}
