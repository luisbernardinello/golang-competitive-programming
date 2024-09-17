package main

import (
	"bufio"
	"fmt"
	"math"
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


func calculaPlacas(num int, consoantes int) int {
	// 26^C * 10^D
	if num == 0 && consoantes == 0 {
		return 0
	}
	return int(math.Pow(26, float64(num)) * math.Pow(10, float64(consoantes)))
}

func solve() {
	casos := nextInt()

	for i := 0; i < casos; i++ {
		numeros := nextInt()
		consoantes := nextInt()

		
		resultado := calculaPlacas(numeros, consoantes)
		out(resultado)
	}
}

func main() {
	solve()
	writer.Flush()
}
