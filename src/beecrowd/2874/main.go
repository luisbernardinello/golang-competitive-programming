package main

import (
	"fmt"
	"strconv"
)

func converteBinarioEmAsc(valores []string) string {
	var resultado string

	for _, valor := range valores {
		numDecimal, _ := strconv.ParseInt(valor, 2, 64)
		char := string(rune(numDecimal))
		resultado += char
	}

	return resultado
}

func main() {
	var casos int

	for {
		_, err := fmt.Scan(&casos)
		if err != nil {
			break
		}

		valores := make([]string, casos)
		for i := 0; i < casos; i++ {
			fmt.Scan(&valores[i])
		}

		frase := converteBinarioEmAsc(valores)
		fmt.Println(frase)
	}
}
