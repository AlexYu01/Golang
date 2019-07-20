package main

import (
	"fmt"
)

func main() {
	var str1, str2 string = "Hello ", "world!"

	fmt.Println(strCat(&str1, &str2))
}
