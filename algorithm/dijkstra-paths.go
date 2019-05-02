package algorithm

import (
	"fmt"

	"github.com/ckybonist/csgo/structure/graph"
)

// IntMax is maximum value of signed integer
const IntMax = int(^uint(0) >> 1)

func minDist(dists []int, visited []bool) int {
	v := 0
	w := IntMax

	for i, visit := range visited {
		if !visit && dists[i] < w {
			w = dists[i]
			v = i
		}
	}

	return v
}

func buildPaths(prev map[int]int, vertices []int) [][]int {
	paths := [][]int{}

	for u := range prev {
		path := []int{}
		_, existed := prev[u]
		for existed {
			path = append(path, u)
			u, existed = prev[u]
		}
		paths = append(paths, path)
	}

	return paths
}

func printPaths(paths [][]int) {
	fmt.Println("\nPaths (right to left):")
	for _, path := range paths {
		for i, v := range path {
			if i == 0 {
				fmt.Printf("{%d}", v)
				continue
			}
			fmt.Printf(" %d", v)
		}
		fmt.Println("")
	}
}

func findShortestPath(g *graph.Graph, src int) ([]int, [][]int) {
	n := len(g.Vertices)
	visited := make([]bool, n)
	// Init distances
	dists := make([]int, n)
	for i := range dists {
		dists[i] = IntMax
	}

	prev := map[int]int{}

	// Source distance is always 0
	dists[src] = 0
	count := 0

	for count < n {
		u := minDist(dists, visited)
		visited[u] = true
		edges := g.Vertices[u]

		// Update distances
		// Edge of u { v, weight }
		for _, edge := range edges {
			// The vertices that we want to update its distances must be
			// unvisited and connected with current vertex
			if !visited[edge.V] && edge.Weight > 0 && dists[u] != IntMax {
				weight := dists[u] + edge.Weight
				if weight < dists[edge.V] {
					dists[edge.V] = weight
					prev[edge.V] = u
				}
			}
		}
		count++
	}

	vertices := []int{}
	for v := range g.Vertices {
		vertices = append(vertices, v)
	}
	paths := buildPaths(prev, vertices)

	return dists, paths
}

// Dijkstra will find shortest path for each vertex
func Dijkstra() {
	data := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6}, {8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}

	graph := graph.NewGraph(data)
	fmt.Println("Graph:")
	graph.Print()
	fmt.Println("")

	dists, paths := findShortestPath(graph, 0)

	fmt.Println("Vertex    Distance from source (0)")
	for v, dist := range dists {
		fmt.Printf("%d         %d\n", v, dist)
	}

	printPaths(paths)
}
