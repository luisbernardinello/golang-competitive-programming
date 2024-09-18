package main

import (
	"fmt"
	"math"
)

type Edge struct {
	a, b    int
	capacity, flow int
}

type Graph struct {
	v    int
	s, t int
	e    []Edge
	g    [][]int
	d    []int
	ptr  []int
	queue []int
}

func NewGraph(vertexes int) *Graph {
	v := vertexes
	s := 0
	t := v - 1
	g := make([][]int, v)
	d := make([]int, v)
	ptr := make([]int, v)
	queue := make([]int, v)
	return &Graph{
		v:     v,
		s:     s,
		t:     t,
		e:     []Edge{},
		g:     g,
		d:     d,
		ptr:   ptr,
		queue: queue,
	}
}

func (g *Graph) AddEdge(v1, v2, capacity int) {
	v1-- 
	v2--
	e1 := Edge{a: v1, b: v2, capacity: capacity, flow: 0}
	e2 := Edge{a: v2, b: v1, capacity: 0, flow: 0}
	g.g[v1] = append(g.g[v1], len(g.e))
	g.e = append(g.e, e1)
	g.g[v2] = append(g.g[v2], len(g.e))
	g.e = append(g.e, e2)
}

func (g *Graph) Dinic() int {
	flow := 0
	for g.BFS() {
		g.ptr = make([]int, g.v)
		for {
			pushed := g.DFS(g.s, math.MaxInt32)
			if pushed == 0 {
				break
			}
			flow += pushed
		}
	}
	return flow
}

func (g *Graph) BFS() bool {
	qhead := 0
	qtail := 0
	g.queue[qtail] = g.s
	qtail++
	g.d = make([]int, g.v)
	for i := range g.d {
		g.d[i] = -1
	}
	g.d[g.s] = 0
	for qhead < qtail {
		v := g.queue[qhead]
		qhead++
		for _, i := range g.g[v] {
			to := g.e[i].b
			if g.d[to] == -1 && g.e[i].flow < g.e[i].capacity {
				g.queue[qtail] = to
				qtail++
				g.d[to] = g.d[v] + 1
			}
		}
	}
	return g.d[g.t] != -1
}

func (g *Graph) DFS(v, minflow int) int {
	if v == g.t {
		return minflow
	}
	for g.ptr[v] < len(g.g[v]) {
		id := g.g[v][g.ptr[v]]
		to := g.e[id].b
		if g.d[to] == g.d[v]+1 {
			pushed := g.DFS(to, min(minflow, g.e[id].capacity-g.e[id].flow))
			if pushed > 0 {
				g.e[id].flow += pushed
				g.e[id^1].flow -= pushed
				return pushed
			}
		}
		g.ptr[v]++
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	g := NewGraph(4)

	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 2)
	g.AddEdge(3, 2, 1)
	g.AddEdge(2, 4, 2)
	g.AddEdge(3, 4, 1)

	fmt.Println("Maximum flow:", g.Dinic())
}
