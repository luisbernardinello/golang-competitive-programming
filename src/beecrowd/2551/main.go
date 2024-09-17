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

func nextFloat64() float64 {
	f, err := strconv.ParseFloat(next(), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func out(a ...interface{}) {
	fmt.Fprintln(writer, a...)
	writer.Flush()
}

func solve() {
	for {
		treinos := nextInt()

		var duracaoTreino, distanciaTreino int
		var recordeAtual float64
		var diasRecorde []int

		for i := 1; i <= treinos; i++ {
			duracaoTreino = nextInt()
			distanciaTreino = nextInt()
			vm := float64(distanciaTreino) / float64(duracaoTreino)

			if i == 1 || vm > recordeAtual {
				recordeAtual = vm
				diasRecorde = append(diasRecorde, i)
			}
		}

		for _, dia := range diasRecorde {
			out(dia)
		}

		if reader.r.Buffered() == 0 {
			break
		}
	}
}

func main() {
	solve()
}
