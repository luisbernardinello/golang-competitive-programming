package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	From, To, Weight int
}

type Graph struct {
	Nodes [][]Edge
}

func NewGraph(size int) *Graph {
	g := &Graph{Nodes: make([][]Edge, size)}
	return g
}

func (g *Graph) AddEdge(edge Edge) {
	g.Nodes[edge.From] = append(g.Nodes[edge.From], edge)
	g.Nodes[edge.To] = append(g.Nodes[edge.To], Edge{From: edge.To, To: edge.From, Weight: edge.Weight})
}

func (g *Graph) GetNeighbours(node int) []Edge {
	return g.Nodes[node]
}

type PathAndDistance struct {
	Distance  int
	Path      []int
	Estimated int
}

func NewPathAndDistance(distance int, path []int, estimated int) *PathAndDistance {
	return &PathAndDistance{Distance: distance, Path: path, Estimated: estimated}
}

type PriorityQueue []*PathAndDistance

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return (pq[i].Distance + pq[i].Estimated) < (pq[j].Distance + pq[j].Estimated)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*PathAndDistance))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func aStar(from, to int, graph *Graph, heuristic []int) *PathAndDistance {
	queue := &PriorityQueue{}
	heap.Init(queue)

	heap.Push(queue, NewPathAndDistance(0, []int{from}, 0))

	for queue.Len() > 0 {
		currentData := heap.Pop(queue).(*PathAndDistance)
		currentPosition := currentData.Path[len(currentData.Path)-1]

		if currentPosition == to {
			return currentData
		}

		for _, edge := range graph.GetNeighbours(currentPosition) {
			if !contains(currentData.Path, edge.To) {
				newPath := append([]int(nil), currentData.Path...)
				newPath = append(newPath, edge.To)
				newDistance := currentData.Distance + edge.Weight
				heap.Push(queue, NewPathAndDistance(newDistance, newPath, heuristic[edge.To]))
			}
		}
	}
	return NewPathAndDistance(-1, nil, -1)
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	heuristic := []int{
		366, 0, 160, 242, 161, 178, 77, 151, 226, 244,
		241, 234, 380, 98, 193, 253, 329, 80, 199, 374,
	}

	graph := NewGraph(20)
	graphData := []int{
		0, 19, 75, 0, 15, 140, 0, 16, 118, 19, 12, 71, 12, 15, 151, 16, 9, 111,
		9, 10, 70, 10, 3, 75, 3, 2, 120, 2, 14, 146, 2, 13, 138, 2, 6, 115,
		15, 14, 80, 15, 5, 99, 14, 13, 97, 5, 1, 211, 13, 1, 101, 6, 1, 160,
		1, 17, 85, 17, 7, 98, 7, 4, 86, 17, 18, 142, 18, 8, 92, 8, 11, 87,
	}

	for i := 0; i < len(graphData); i += 3 {
		graph.AddEdge(Edge{From: graphData[i], To: graphData[i+1], Weight: graphData[i+2]})
	}

	solution := aStar(3, 1, graph, heuristic)
	printSolution(solution)
}

func printSolution(p *PathAndDistance) {
	if p.Path != nil {
		fmt.Printf("Optimal path: %v, distance: %d\n", p.Path, p.Distance)
	} else {
		fmt.Println("There is no path available to connect the points")
	}
}
