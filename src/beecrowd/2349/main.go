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

func out(a ...interface{}) {
	fmt.Fprintln(writer, a...)
	writer.Flush()
}

func solve() {
	estacoes := nextInt()
	comandos := nextInt()
	estacaoProx := nextInt()
	posicaoAtual := 1
	visitas := 0

	comandoSequencia := make([]int, comandos)
	for i := 0; i < comandos; i++ {
		comandoSequencia[i] = nextInt()
	}

	if posicaoAtual == estacaoProx {
		visitas++
	}

	for _, comando := range comandoSequencia {
		if comando == 1 {
			posicaoAtual++
			if posicaoAtual > estacoes {
				posicaoAtual = 1
			}
		} else if comando == -1 {
			posicaoAtual--
			if posicaoAtual < 1 {
				posicaoAtual = estacoes
			}
		}

		if posicaoAtual == estacaoProx {
			visitas++
		}
	}

	out(visitas)
}

func main() {
	solve()
}
