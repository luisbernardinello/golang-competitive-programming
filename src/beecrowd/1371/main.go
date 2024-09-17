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

func calcSequencia(n int) []int {
	var portasAbertas []int
	for i := 1; i*i <= n; i++ {
		portasAbertas = append(portasAbertas, i*i)
	}
	return portasAbertas
}

func solve() {
	for {
		cases := nextInt()
		if cases == 0 {
			break
		}

		resultado := calcSequencia(cases)
		for i, val := range resultado {
			if i > 0 {
				writer.WriteString(" ")
			}
			writer.WriteString(strconv.Itoa(val))
		}
		writer.WriteString("\n")
		writer.Flush() 
	}
}

func main() {
	solve()
}
