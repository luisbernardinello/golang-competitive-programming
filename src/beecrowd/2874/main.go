package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bufReader struct {
	r   *bufio.Reader
	buf []byte
	i   int
}

var reader = &bufReader{
	bufio.NewReader(os.Stdin),
	make([]byte, 0),
	0,
}

var writer = bufio.NewWriter(os.Stdout)

func (r *bufReader) readLine() {
	if r.i < len(r.buf) {
		return
	}
	r.buf = make([]byte, 0)
	r.i = 0
	for {
		line, isPrefix, err := r.r.ReadLine()
		if err != nil {
			panic(err)
		}
		r.buf = append(r.buf, line...)
		if !isPrefix {
			break
		}
	}
}

func (r *bufReader) next() string {
	r.readLine()
	from := r.i
	for ; r.i < len(r.buf); r.i++ {
		if r.buf[r.i] == ' ' {
			break
		}
	}
	s := string(r.buf[from:r.i])
	r.i++
	return s
}

func (r *bufReader) nextLine() string {
	r.readLine()
	s := string(r.buf[r.i:])
	r.i = len(r.buf)
	return s
}

func next() string {
	return reader.next()
}

func nextInt() int {
	i, err := strconv.Atoi(next())
	if err != nil {
		panic(err)
	}
	return i
}

func nextLine() string {
	return reader.nextLine()
}

func out(a ...interface{}) {
	fmt.Fprintln(writer, a...)
	writer.Flush()
}

// Função para converter binário em ASCII
func converteBinarioEmAsc(valores []string) string {
	var resultado string

	for _, valor := range valores {
		numDecimal, _ := strconv.ParseInt(valor, 2, 64)
		char := string(rune(numDecimal))
		resultado += char
	}

	return resultado
}

func solve() {
	// Leia os dados até o fim do arquivo (EOF)
	for {
		var casosStr string
		if _, err := fmt.Fscanf(reader.r, "%s\n", &casosStr); err != nil {
			break // EOF
		}

		casos, _ := strconv.Atoi(casosStr)
		valores := make([]string, casos)

		for i := 0; i < casos; i++ {
			valores[i] = next()
		}

		// Converta os valores binários em ASCII e imprima
		frase := converteBinarioEmAsc(valores)
		out(frase)
	}
}

func main() {
	solve()
	//writer.Flush()
}
