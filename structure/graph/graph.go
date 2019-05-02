package graph

import (
	"fmt"
	"sort"
)

type Edge struct {
	V      int
	Weight int
}
type edges []Edge
type vertices map[int]edges
type Graph struct {
	Vertices vertices
}

func NewGraph(matrix [][]int) *Graph {
	vs := vertices{}

	for u, vlist := range matrix {
		edges := edges{}
		for v, weight := range vlist {
			edges = append(edges, Edge{v, weight})
		}
		vs[u] = edges
	}

	return &Graph{vs}
}

func (g *Graph) Print() {
	var keys []int
	for k := range g.Vertices {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		edges := g.Vertices[v]
		fmt.Printf("%d => %v\n", v, edges)
	}
}
