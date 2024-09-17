// This template is made by Luis Bernardinello  <luis.bernardinello@gmail.com>
// GitHub : https://github.com/luisbernardinello

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type IO struct {
	r   *bufio.Reader
	w   *bufio.Writer
	buf []string
	i   int
}

var io = &IO{
	r: bufio.NewReaderSize(os.Stdin, 1<<20), // 1MB buffer
	w: bufio.NewWriterSize(os.Stdout, 1<<20), 
}

func (io *IO) loadBuffer() {
	if io.i < len(io.buf) {
		return
	}
	io.i, io.buf = 0, nil
	line, err := io.r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	io.buf = strings.Fields(line)
}

func (io *IO) next() string {
	io.loadBuffer()
	token := io.buf[io.i]
	io.i++
	return token
}

func (io *IO) nextInt() int {
	val, err := strconv.Atoi(io.next())
	if err != nil {
		panic(err)
	}
	return val
}

func (io *IO) nextFloat() float64 {
	val, err := strconv.ParseFloat(io.next(), 64)
	if err != nil {
		panic(err)
	}
	return val
}

func (io *IO) print(a ...interface{}) {
	for _, v := range a {
		switch v := v.(type) {
		case int:
			io.w.WriteString(strconv.Itoa(v))
		case float64:
			io.w.WriteString(strconv.FormatFloat(v, 'f', 6, 64))
		case string:
			io.w.WriteString(v)
		default:
			panic("unsupported type")
		}
		io.w.WriteByte(' ')
	}
	io.w.WriteByte('\n')
}

func (io *IO) flush() {
	io.w.Flush()
}

// AUX FUNCTIONS

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func pow(x, y int) int {
	result := 1
	for y > 0 {
		if y&1 != 0 {
			result *= x
		}
		x *= x
		y >>= 1
	}
	return result
}

// YOUR CODE HERE

func solve() {
	// use cases
	print("enter with two integers and two floats:\n")
	n := io.nextInt()
	m := io.nextInt()
	f := io.nextFloat()
	g := io.nextFloat()
	print("the sum is:\n")
	io.print("integer + integer =", n + m)
	io.print("float + float =", f + g)
}

func main() {
	solve()
	io.flush()
}
