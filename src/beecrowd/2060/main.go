package main

import (
	"fmt"
)

func contaMultiplos(valores []int) {
	count2, count3, count4, count5 := 0, 0, 0, 0

	for _, valor := range valores {
		if valor%2 == 0 {
			count2++
		}
		if valor%3 == 0 {
			count3++
		}
		if valor%4 == 0 {
			count4++
		}
		if valor%5 == 0 {
			count5++
		}
	}

	fmt.Printf("%d Multiplo(s) de 2\n", count2)
	fmt.Printf("%d Multiplo(s) de 3\n", count3)
	fmt.Printf("%d Multiplo(s) de 4\n", count4)
	fmt.Printf("%d Multiplo(s) de 5\n", count5)
}

func main() {
	var casos int
	fmt.Scan(&casos)

	valores := make([]int, casos)
	for i := 0; i < casos; i++ {
		fmt.Scan(&valores[i])
	}

	contaMultiplos(valores)
}
