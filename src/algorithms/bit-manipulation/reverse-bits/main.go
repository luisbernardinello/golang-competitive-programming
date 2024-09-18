package main

import "fmt"

// reverte os bits de um número 32 bits
func reverseBits(n uint32) uint32 {
	var retval uint32 = 0

	for i := 0; i < 32; i++ {
		retval <<= 1
		retval |= n & 1
		n >>= 1
	}

	return retval
}

func main() {
	// (Input)  12345678:    00000000 10111100 01100001 01001110
	// (Output) 1921400064:  01110010 10000110 00111101 00000000
	fmt.Println(reverseBits(12345678))
}
