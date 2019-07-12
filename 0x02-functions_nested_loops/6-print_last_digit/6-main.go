package main

import (
	"fmt"
)

func main() {
	var num1 int = 1234
	var num2 int = 10

	fmt.Printf("The last digit of %d is %d\n", num1, lastDigit(num1))
	fmt.Printf("The last digit of %d is %d\n", num2, lastDigit(num2))
}
