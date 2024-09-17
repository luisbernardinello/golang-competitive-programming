package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func criaContador() (func(int), func() []int, func() map[int]int) {
	valores := []int{}
	contagem := map[int]int{}

	addValor := func(valor int) {
		valores = append(valores, valor)
		sort.Ints(valores)
		contagem[valor]++
	}


	getValores := func() []int {
		return valores
	}

	getContagem := func() map[int]int {
		return contagem
	}

	return addValor, getValores, getContagem
}

func solve() {
	addValor, getValores, getContagem := criaContador()

	casos := nextInt()

	for i := 0; i < casos; i++ {
		valor := nextInt()
		addValor(valor)
	}

	writer.WriteString("Valores ordenados: ")
	valores := getValores()
	for i, val := range valores {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(strconv.Itoa(val))
	}
	writer.WriteString("\n")
	writer.Flush()

	
	for numero, contagem := range getContagem() {
		out(fmt.Sprintf("Número %d aparece %d vezes", numero, contagem))
	}
}

func main() {
	solve()
}
