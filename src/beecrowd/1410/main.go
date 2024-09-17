package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func calcImpedimento(distAtac []int, distDef []int) string {
	sort.Ints(distDef)
	sort.Ints(distAtac)

	if distAtac[0] < distDef[1] {
		return "Y"
	}
	return "N"
}

func solve() {
	for {
		jogadores := nextInt()
		defensores := nextInt()

		if jogadores == 0 && defensores == 0 {
			break
		}

		distanciaAtacantes := make([]int, jogadores)
		for i := 0; i < jogadores; i++ {
			distanciaAtacantes[i] = nextInt()
		}

		distanciaDefensores := make([]int, defensores)
		for i := 0; i < defensores; i++ {
			distanciaDefensores[i] = nextInt()
		}

		impedido := calcImpedimento(distanciaAtacantes, distanciaDefensores)
		out(impedido)
	}
}

func main() {
	solve()
}
