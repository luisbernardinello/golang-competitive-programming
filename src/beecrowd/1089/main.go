package main

import (
	"bufio"
	"container/ring"
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

func calcPico(r *ring.Ring) int {
	picos := 0
	current := r

	for i := 0; i < r.Len(); i++ {
		currentValue := current.Value.(int)
		prevValue := current.Prev().Value.(int)
		nextValue := current.Next().Value.(int)

		if (currentValue > prevValue && currentValue > nextValue) ||
			(currentValue < prevValue && currentValue < nextValue) {
			picos++
		}

		current = current.Next()
	}

	return picos
}

func solve() {
	for {
		cases := nextInt()
		if cases == 0 {
			break
		}

		r := ring.New(cases)

		for i := 0; i < cases; i++ {
			value := nextInt()
			r.Value = value
			r = r.Next()
		}

		resultado := calcPico(r)
		out(resultado)
	}
}

func main() {
	solve()
}
