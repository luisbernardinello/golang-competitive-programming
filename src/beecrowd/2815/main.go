package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (r *bufReader) nextLine() string {
	r.readLine()
	s := string(r.buf[r.i:])
	r.i = len(r.buf)
	return s
}

var writer = bufio.NewWriter(os.Stdout)

func next() string {
	return reader.next()
}

func nextLine() string {
	return reader.nextLine()
}

func out(a ...interface{}) {
	fmt.Fprintln(writer, a...)
	writer.Flush()
}

func solve() {
	line := nextLine()

	words := strings.Split(line, " ")

	var correctedWords []string

	for _, word := range words {
		if len(word) > 2 && word[:2] == word[2:4] {
			correctedWords = append(correctedWords, word[2:])
		} else {
			correctedWords = append(correctedWords, word)
		}
	}

	out(strings.Join(correctedWords, " "))
}

func main() {
	solve()
}

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (r *bufReader) nextLine() string {
	r.readLine()
	s := string(r.buf[r.i:])
	r.i = len(r.buf)
	return s
}

var writer = bufio.NewWriter(os.Stdout)

func next() string {
	return reader.next()
}

func nextLine() string {
	return reader.nextLine()
}

func out(a ...interface{}) {
	fmt.Fprintln(writer, a...)
	writer.Flush()
}

func solve() {
	line := nextLine()

	words := strings.Split(line, " ")

	var correctedWords []string

	for _, word := range words {
		if len(word) > 2 && word[:2] == word[2:4] {
			correctedWords = append(correctedWords, word[2:])
		} else {
			correctedWords = append(correctedWords, word)
		}
	}

	out(strings.Join(correctedWords, " "))
}

func main() {
	solve()
}
