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

func calcZeroOuUm(alice int, beto int, clara int) string {
	if alice != beto && alice != clara {
		return "A"
	} else if beto != alice && beto != clara {
		return "B"
	} else if clara != alice && clara != beto {
		return "C"
	} else {
		return "*"
	}
}

func solve() {
	for {
		// Lê os valores de A, B e C
		aStr := next()
		if aStr == "" {
			break
		}
		alice := nextInt()
		beto := nextInt()
		clara := nextInt()

		resultado := calcZeroOuUm(alice, beto, clara)


		out(resultado)
	}
}

func main() {
	solve()
}
