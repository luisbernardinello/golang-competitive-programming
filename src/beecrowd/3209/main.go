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

func out(format string, a ...interface{}) {
	fmt.Fprintf(writer, format+"\n", a...)
	writer.Flush()
}

func calcTomadasFiltros(vet []int) int {
	totalTomadas := 0
	for _, tomadas := range vet {
		totalTomadas += tomadas
	}
	return totalTomadas - (len(vet) - 1)
}

func solve() {
	cases := nextInt()

	for cases > 0 {
		cases--

		filtros := nextInt()
		tomadas := make([]int, filtros)
		for i := 0; i < filtros; i++ {
			tomadas[i] = nextInt()
		}

		resultado := calcTomadasFiltros(tomadas)
		out("%d", resultado)
	}
}

func main() {
	solve()
}
