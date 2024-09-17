package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		casos, _ := strconv.Atoi(line)

		if casos == 0 {
			break
		}

		ordemCorreta := make([]int, casos)
		for i := 0; i < casos; i++ {
			ordemCorreta[i] = i + 1
		}

		tentativas := 0
		for {
			tentativas++
			line, _ := reader.ReadString('\n')
			valoresStr := strings.Fields(line)
			valores := make([]int, casos)
			isEqual := true

			for i := 0; i < casos; i++ {
				valores[i], _ = strconv.Atoi(valoresStr[i])
				if valores[i] != ordemCorreta[i] {
					isEqual = false
				}
			}

			if isEqual {
				break
			}
		}

		fmt.Fprintln(writer, tentativas)
	}
}
