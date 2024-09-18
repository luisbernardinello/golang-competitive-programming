package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100
const INF = 2000000000 

func pushExcess(u, v int, f *[N + 2][N + 2]int, c *[N + 2][N + 2]int, xf *[N + 2]int) {
	df := min(xf[u], c[u][v]-f[u][v])
	f[u][v] += df
	f[v][u] -= df
	xf[u] -= df
	xf[v] += df
}

func relabelNode(u int, c *[N + 2][N + 2]int, f *[N + 2][N + 2]int, ht *[N + 2]int) {
	minHt := 2 * N
	for v := 0; v < N; v++ {
		if c[u][v] > f[u][v] {
			minHt = min(minHt, ht[v])
		}
	}
	ht[u] = minHt + 1
}

func discharge(u, nodes int, f *[N + 2][N + 2]int, c *[N + 2][N + 2]int, xf *[N + 2]int, ht *[N + 2]int, nxt *[N + 2]int) {
	for xf[u] > 0 {
		v := nxt[u]
		if v < nodes {
			if c[u][v] > f[u][v] && ht[u] > ht[v] {
				pushExcess(u, v, f, c, xf)
			} else {
				nxt[u]++
			}
		} else {
			relabelNode(u, c, f, ht)
			nxt[u] = 0
		}
	}
}

func goldbergTarjan(src, sink, nodes int, f *[N + 2][N + 2]int, c *[N + 2][N + 2]int, xf *[N + 2]int, ht *[N + 2]int, nxt *[N + 2]int) int {
	var list [N]int
	maxFlow := 0

	for i := 0; i < nodes; i++ {
		if i != src && i != sink {
			list[i] = i
		}
	}

	ht[src] = nodes
	xf[src] = INF

	for i := 0; i < nodes; i++ {
		if i != src {
			pushExcess(src, i, f, c, xf)
		}
	}

	index := 0
	for index < nodes-2 {
		u := list[index]
		prevHt := ht[u]
		discharge(u, nodes, f, c, xf, ht, nxt)
		if ht[u] > prevHt {
			for j := index; j > 0; j-- {
				list[j] = list[j-1]
			}
			list[0] = u
		} else {
			index++
		}
	}

	for v := 0; v < nodes; v++ {
		maxFlow += f[src][v]
	}

	return maxFlow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o número de nós: ")
	nodesStr, _ := reader.ReadString('\n')
	nodesStr = strings.TrimSpace(nodesStr)
	nodes, _ := strconv.Atoi(nodesStr)

	var c, f [N + 2][N + 2]int

	for i := 0; i < nodes; i++ {
		fmt.Printf("Digite as capacidades da linha %d: ", i)
		input, _ := reader.ReadString('\n')
		capacities := strings.Fields(input)
		for j := 0; j < nodes; j++ {
			c[i][j], _ = strconv.Atoi(capacities[j])
			f[i][j] = 0
		}
	}

	var xf [N + 2]int
	var ht [N + 2]int
	var nxt [N + 2]int

	maxFlow := goldbergTarjan(0, nodes-1, nodes, &f, &c, &xf, &ht, &nxt)
	fmt.Printf("Fluxo máximo: %d\n", maxFlow)
}
