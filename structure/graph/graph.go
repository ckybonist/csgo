package graph

type Edge struct {
	U int
	V int
}
type edges []Edge
type vertices map[int]edges
type Graph struct {
	Vertices vertices
}

func NewGraph(matrix [][]int) *Graph {
	var vs vertices

	for u, vlist := range matrix {
		var edges edges
		for v, weight := range vlist {
			edges = append(edges, Edge{v, weight})
		}
		vs[u] = edges
	}

	return &Graph{vs}
}
