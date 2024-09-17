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

func out(a ...interface{}) {
	fmt.Fprintln(writer, a...)
}

func nextInt64() int64 {
	i, err := strconv.ParseInt(reader.next(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func nextInt() int {
	return int(nextInt64())
}


func contaMultiplos(valores []int) {
	count2, count3, count4, count5 := 0, 0, 0, 0

	for _, valor := range valores {
		if valor%2 == 0 {
			count2++
		}
		if valor%3 == 0 {
			count3++
		}
		if valor%4 == 0 {
			count4++
		}
		if valor%5 == 0 {
			count5++
		}
	}

	out(count2, "Multiplo(s) de 2")
	out(count3, "Multiplo(s) de 3")
	out(count4, "Multiplo(s) de 4")
	out(count5, "Multiplo(s) de 5")
}

func solve() {
	casos := nextInt()

	valores := make([]int, casos)

	for i := 0; i < casos; i++ {
		valores[i] = nextInt()
	}

	contaMultiplos(valores)
}

func main() {

	solve()
	writer.Flush()
}