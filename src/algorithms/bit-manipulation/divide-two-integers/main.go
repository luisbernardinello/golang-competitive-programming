package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func para dividir sem usar multiplicação, mod ou divisão
func divide(a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não é permitida")
	}


	// det o sinal do resultado
	sign := int32(1)
	if (a < 0) != (b < 0) {
		sign = -1
	}

	// valores absolutos
	ua := int64(abs(a))
	ub := int64(abs(b))

	// divide utilizando subtracao e deslocamento de bits
	var result int32 = 0
	for ua >= ub {
		temp, multiple := ub, int32(1)
		for ua >= (temp << 1) {
			temp <<= 1
			multiple <<= 1
		}
		ua -= temp
		result += multiple
	}

	return result * sign, nil
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Entre com dois numeros separados por espaço:")
	input, _ := reader.ReadString('\n')
	numbers := strings.Fields(input)
	if len(numbers) != 2 {
		fmt.Println("atencao, insira exatamente dois numeros.")
		return
	}

	a, err1 := strconv.ParseInt(numbers[0], 10, 32)
	b, err2 := strconv.ParseInt(numbers[1], 10, 32)
	if err1 != nil || err2 != nil {
		fmt.Println("numeros invalidos, insira numeros validos.")
		return
	}

	result, err := divide(int32(a), int32(b))
	if err != nil {
		fmt.Println("gerou erro:", err)
		return
	}

	fmt.Printf("a divisão é: %d\n", result)
}
