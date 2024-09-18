package main

import "fmt"

// calculo de potencia usando divisão e conquista O(logN) tempo e O(logN) espaço
func fastPower(a, b int) int {
	if b == 0 {
		return 1
	}
	smallerValue := fastPower(a, b/2)
	result := smallerValue * smallerValue
	if b%2 == 1 {
		result *= a
	}
	return result
}

// calculo de potencia usando BitMask O(logN) tempo e O(1) espaço
func fastPowerBitmask(a, b int) int {
	res := 1
	for b > 0 {
		lastBit := b & 1
		if lastBit == 1 {
			res *= a
		}
		a *= a
		b >>= 1 // desloca para direita removendo o bit menos significativo
	}
	return res
}

func main() {
	fmt.Println(fastPower(3, 5))
	fmt.Println(fastPowerBitmask(3, 5))
}
