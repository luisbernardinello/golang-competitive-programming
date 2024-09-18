package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func copySetBitsInRange(a *uint32, b uint32, l uint32, r uint32) {

	if l < 1 || r > 32 {
		return
	}

	// para um dado intervalo
	for i := l; i <= r; i++ {
		// cria máscara ( número cujo unico bit definido está na posição i)
		mask := uint32(1) << (i - 1)

		// se o bit i estiver definido em b, definir o bit i em a 
		if b&mask != 0 {
			*a |= mask
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Insira o primeiro número: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	tempA, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("numero invalido, insira um número válido.")
		return
	}
	a := uint32(tempA)

	fmt.Print("Insira o segundo número: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	tempB, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("numero invalido, insira um número válido.")
		return
	}
	b := uint32(tempB)

	fmt.Print("Insira o intervalo (l e r): ")
	input, _ = reader.ReadString('\n')
	rangeValues := strings.Fields(input)
	if len(rangeValues) != 2 {
		fmt.Println("numero invalido, insira dois valores para o intervalo.")
		return
	}
	l, err := strconv.ParseUint(rangeValues[0], 10, 32)
	if err != nil {
		fmt.Println("numero invalido, insira um número válido.")
		return
	}
	r, err := strconv.ParseUint(rangeValues[1], 10, 32)
	if err != nil {
		fmt.Println("numero invalido, insira um número válido.")
		return
	}

	copySetBitsInRange(&a, b, uint32(l), uint32(r))

	fmt.Printf("valor modificado do do primeiro numero depois de copiar bits setados no intervalo - %d\n", a)
}
