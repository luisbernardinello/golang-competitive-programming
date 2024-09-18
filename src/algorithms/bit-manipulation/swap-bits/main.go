package main

import "fmt"

// troca `n` bits em `x` nas posições `p1` e `p2`
func swapBits(x, p1, p2, n uint32) uint32 {
	//joga todos os bits do primeiro conjunto para o lado direito
	set1 := (x >> p1) & ((1 << n) - 1)

	// joga todos os bits do segundo conjunto para o lado direito
	set2 := (x >> p2) & ((1 << n) - 1)

	// XOR os dois conjuntos
	xor := set1 ^ set2

	// Coloca os bits XOR de volta nas posicoes originais
	xor = (xor << p1) | (xor << p2)

	// XOR o resultado com o num original para os dois conjuntos sejam trocados
	result := x ^ xor

	return result
}

func main() {

	res := swapBits(28, 0, 3, 2)
	fmt.Println("Resultado =", res)
}
