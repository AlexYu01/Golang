package main

// revArray Takes a int slice and reverses the underlying array's content
func revArray(a []int) {
	var midP, back, front uint

	midP = uint(len(a)) / 2
	back = uint(len(a)) - 1
	front = 0

	for ; front < midP; front++ {
		a[front], a[back] = a[back], a[front]
		back--
	}
}
