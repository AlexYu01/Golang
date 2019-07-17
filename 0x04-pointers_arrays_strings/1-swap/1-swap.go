package main

// swapInt Takes two int pointers and swap the integer values in the memory location
func swapInt(m, n *int) {
	*m, *n = *n, *m

}
