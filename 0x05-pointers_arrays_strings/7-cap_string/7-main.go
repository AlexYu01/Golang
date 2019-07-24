package main

import (
	"fmt"
)

func main() {
	src := "Expect the best. Prepare for the worst. Capitalize on what comes.\nhello world! hello-world 0123456hello world\thello world.hello world\n"

	fmt.Print(src)
	fmt.Print(capString(src))
}
