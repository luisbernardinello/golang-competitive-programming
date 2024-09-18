package main

import (
	"container/list"
	"fmt"
	"math"
)

func edmondsKarp(adjMatrix, capacityMatrix [][]int, source, sink int) int {
	n := len(capacityMatrix)
	flow := 0
	flowMatrix := make([][]int, n)
	for i := range flowMatrix {
		flowMatrix[i] = make([]int, n)
	}

	for {
		p := make([]int, n)
		for i := range p {
			p[i] = -1
		}
		p[source] = -2 
		m := make([]int, n)
		m[source] = math.MaxInt32

		bfsQueue := list.New()
		bfsQueue.PushBack(source)

		pathFlow, p := bfsEK(adjMatrix, capacityMatrix, sink, flowMatrix, p, m, bfsQueue)
		if pathFlow == 0 {
			break 
		}

		flow += pathFlow
		v := sink
		for v != source {
			u := p[v]
			flowMatrix[u][v] += pathFlow
			flowMatrix[v][u] -= pathFlow
			v = u
		}
	}

	return flow
}

func bfsEK(adjMatrix, capacityMatrix [][]int, sink int, flowMatrix [][]int, p, m []int, bfsQueue *list.List) (int, []int) {
	for bfsQueue.Len() > 0 {
		u := bfsQueue.Remove(bfsQueue.Front()).(int)
		for _, v := range adjMatrix[u] {
			if capacityMatrix[u][v]-flowMatrix[u][v] > 0 && p[v] == -1 {
				p[v] = u 
				m[v] = min(m[u], capacityMatrix[u][v]-flowMatrix[u][v])
				if v != sink {
					bfsQueue.PushBack(v)
				} else {
					return m[sink], p 
				}
			}
		}
	}
	return 0, p 
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	adjMatrix := [][]int{
		{1, 2}, // 0 -> 1, 0 -> 2
		{2, 3}, // 1 -> 2, 1 -> 3
		{3},    // 2 -> 3
		{},     // 3 (sink)
	}
	capacityMatrix := [][]int{
		{0, 1000000, 1000000, 0},
		{0, 0, 1, 1000000},
		{0, 0, 0, 1000000},
		{0, 0, 0, 0},
	}
	source := 0
	sink := 3

	fmt.Println("Maximum flow:", edmondsKarp(adjMatrix, capacityMatrix, source, sink))
}
