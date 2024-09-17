package main

import (
	"fmt"
	"math"
)

func calculaPlacas(num int, consoantes int) int {
	// 26^C * 10^D
	if num == 0 && consoantes == 0 {
		return 0
	}
	return int(math.Pow(26, float64(num)) * math.Pow(10, float64(consoantes)))
}

func main() {
	var casos, numeros, consoantes int
	fmt.Scan(&casos)

	for i := 0; i < casos; i++ {

		fmt.Scan(&numeros)
		fmt.Scan(&consoantes)
		resultado := calculaPlacas(numeros, consoantes)
		fmt.Println(resultado)

	}
}
