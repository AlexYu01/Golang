package main

import (
	"fmt"
)

func main() {
	var src1, src2 = "Today is ", "Friday!"

	fmt.Println(strNCat(&src1, &src2, 4))
}
