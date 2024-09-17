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
	writer.Flush()
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

func solve() {
	for {
		casos := nextInt()

		if casos == 0 {
			break
		}

		ordemCorreta := make([]int, casos)
		for i := 0; i < casos; i++ {
			ordemCorreta[i] = i + 1
		}

		tentativas := 0
		for {
			tentativas++

			valores := make([]int, casos)
			isEqual := true

			for i := 0; i < casos; i++ {
				valores[i] = nextInt()
				if valores[i] != ordemCorreta[i] {
					isEqual = false
				}
			}


			if isEqual {
				break
			}
		}

		out(tentativas)
	}
}

func main() {

	solve()
	//writer.Flush()
}
