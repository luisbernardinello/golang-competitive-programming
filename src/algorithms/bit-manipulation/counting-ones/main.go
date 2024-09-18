package main

import "fmt"

// CONTA os bits '1'
func countOnes(x uint32) uint16 {
	var numBits uint16 = 0
	for x != 0 {
		numBits += uint16(x & 1) // máscara de um bit
		x >>= 1
	}
	return numBits
}

func main() {
	number := uint32(110) 
	result := countOnes(number)
	fmt.Printf("Resultado: %d\n", result)
}
