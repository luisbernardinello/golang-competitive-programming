package main

import (
	"container/list"
	"fmt"
)

const (
	MXN = 210
	INF = 1000000010
)

type Edge struct {
	u, v     int
	cap, flow int64
}

type Dinic struct {
	n   int
	e   []Edge
	g   [][]int
	d   []int
	pt  []int
}

func NewDinic(n int) *Dinic {
	g := make([][]int, n)
	for i := range g {
		g[i] = []int{}
	}
	d := make([]int, n)
	pt := make([]int, n)
	return &Dinic{n: n, e: []Edge{}, g: g, d: d, pt: pt}
}

func (d *Dinic) AddEdge(u, v int, cap int64) {
	if u != v {
		d.e = append(d.e, Edge{u, v, cap, 0})
		d.g[u] = append(d.g[u], len(d.e)-1)
		d.e = append(d.e, Edge{v, u, 0, 0})
		d.g[v] = append(d.g[v], len(d.e)-1)
	}
}

func (d *Dinic) BFS(s, t int) bool {
	q := list.New()
	for i := range d.d {
		d.d[i] = d.n + 1
	}
	d.d[s] = 0
	q.PushBack(s)
	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		if u == t {
			break
		}
		for _, k := range d.g[u] {
			e := &d.e[k]
			if e.flow < e.cap && d.d[e.v] > d.d[e.u]+1 {
				d.d[e.v] = d.d[e.u] + 1
				q.PushBack(e.v)
			}
		}
	}
	return d.d[t] != d.n+1
}

func (d *Dinic) DFS(u, t int, flow int64) int64 {
	if u == t || flow == 0 {
		return flow
	}
	for i := d.pt[u]; i < len(d.g[u]); i++ {
		k := d.g[u][i]
		e := &d.e[k]
		oe := &d.e[k^1]
		if d.d[e.v] == d.d[e.u]+1 {
			amt := e.cap - e.flow
			pushed := d.DFS(e.v, t, min(flow, amt))
			if pushed > 0 {
				e.flow += pushed
				oe.flow -= pushed
				return pushed
			}
		}
		d.pt[u]++
	}
	return 0
}

func (d *Dinic) MaxFlow(s, t int) int64 {
	total := int64(0)
	for d.BFS(s, t) {
		for i := range d.pt {
			d.pt[i] = 0
		}
		for {
			flow := d.DFS(s, t, INF)
			if flow == 0 {
				break
			}
			total += flow
		}
	}
	return total
}

type GomoryHu struct {
	ok     [MXN]int
	cap    [MXN][MXN]int
	answer [MXN][MXN]int
	parent [MXN]int
	n      int
	flow   *Dinic
}

func NewGomoryHu(n int) *GomoryHu {
	g := &GomoryHu{
		n: n,
		flow: NewDinic(n),
	}
	for i := 0; i < MXN; i++ {
		for j := 0; j < MXN; j++ {
			g.answer[i][j] = INF
		}
	}
	return g
}

func (gh *GomoryHu) AddEdge(u, v, c int) {
	gh.cap[u][v] += c
}

func (gh *GomoryHu) Calc() {
	for i := 1; i < gh.n; i++ {
		gh.flow = NewDinic(gh.n)
		for u := 0; u < gh.n; u++ {
			for v := 0; v < gh.n; v++ {
				if gh.cap[u][v] > 0 {
					gh.flow.AddEdge(u, v, int64(gh.cap[u][v]))
				}
			}
		}

		f := int(gh.flow.MaxFlow(i, gh.parent[i]))
		gh.BFS(i)

		for j := i + 1; j < gh.n; j++ {
			if gh.ok[j] != 0 && gh.parent[j] == gh.parent[i] {
				gh.parent[j] = i
			}
		}

		gh.answer[i][gh.parent[i]] = f
		gh.answer[gh.parent[i]][i] = f 

		for j := 0; j < i; j++ {
			gh.answer[i][j] = minInt(gh.answer[i][j], gh.answer[gh.parent[i]][j])
			gh.answer[j][i] = gh.answer[i][j] 
		}
	}
}


func (gh *GomoryHu) BFS(start int) {
	for i := range gh.ok {
		gh.ok[i] = 0
	}
	q := list.New()
	q.PushBack(start)
	gh.ok[start] = 1
	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		for _, xid := range gh.flow.g[u] {
			id := xid
			v := gh.flow.e[id].v
			fl := gh.flow.e[id].flow
			cap := gh.flow.e[id].cap
			if gh.ok[v] == 0 && fl < cap {  
				gh.ok[v] = 1
				q.PushBack(v)
			}
		}
	}
}

func (gh *GomoryHu) GetTree() [MXN][MXN]int {
	return gh.answer
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	flow := NewDinic(6)
	flow.AddEdge(0, 1, 16)
	flow.AddEdge(0, 2, 13)
	flow.AddEdge(1, 2, 10)
	flow.AddEdge(1, 3, 12)
	flow.AddEdge(2, 1, 4)
	flow.AddEdge(2, 4, 14)
	flow.AddEdge(3, 2, 9)
	flow.AddEdge(3, 5, 20)
	flow.AddEdge(4, 3, 7)
	flow.AddEdge(4, 5, 4)

	maxFlow := flow.MaxFlow(0, 5)
	fmt.Printf("Max flow: %d\n", maxFlow)

	gomoryHu := NewGomoryHu(6)
	gomoryHu.AddEdge(0, 1, 16)
	gomoryHu.AddEdge(0, 2, 13)
	gomoryHu.AddEdge(1, 2, 10)
	gomoryHu.AddEdge(1, 3, 12)
	gomoryHu.AddEdge(2, 1, 4)
	gomoryHu.AddEdge(2, 4, 14)
	gomoryHu.AddEdge(3, 2, 9)
	gomoryHu.AddEdge(3, 5, 20)
	gomoryHu.AddEdge(4, 3, 7)
	gomoryHu.AddEdge(4, 5, 4)

	gomoryHu.Calc()
	tree := gomoryHu.GetTree()

	fmt.Println("Gomory-Hu Tree:")
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("%d ", tree[i][j])
		}
		fmt.Println()
	}
}
