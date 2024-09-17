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

func nextFloat64() float64 {
	f, err := strconv.ParseFloat(next(), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func out(format string, a ...interface{}) {
	fmt.Fprintf(writer, format+"\n", a...)
	writer.Flush()
}


func calcSoma(matriz [12][12]float64) float64 {
	total := 0.0
	for i, row := range matriz {
		for j := range row {
			if j > i {
				total += matriz[i][j]
			}
		}
	}
	return total
}

func calcMedia(matriz [12][12]float64) float64 {
	total := calcSoma(matriz)
	count := (12 * 11) / 2
	return total / float64(count)
}

func solve() {
	var m [12][12]float64
	escolha := next()

	for i := range m {
		for j := range m[i] {
			m[i][j] = nextFloat64()
		}
	}

	if escolha == "S" {
		resultado := calcSoma(m)
		out("%.1f", resultado)
	} else if escolha == "M" {
		resultado := calcMedia(m)
		out("%.1f", resultado)
	}
}

func main() {
	solve()
}
