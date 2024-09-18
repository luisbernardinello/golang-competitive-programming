package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func clearIthBit(n int32, i uint32) int32 {
	return n &^ (1 << i)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("insira o numero: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		fmt.Println("numero invalido, insira um numero válido.")
		return
	}

	fmt.Print("insira a posição do bit que sera limpo (indexado em 0): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	i, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("por favor, insira uma posição válida.")
		return
	}

	result := clearIthBit(int32(n), uint32(i))

	fmt.Printf("numero apos limpar o %dº bit: %d\n", i, result)
}
