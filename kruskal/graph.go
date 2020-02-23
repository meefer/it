package kruskal

import (
	"fmt"
	"math/rand"
)

// Node represents a node of a graph
type Node struct {
	x, y float64
}

// Graph represents a graph math structure
type Graph struct {
	nodes []*Node
	edges [][]int
}

func (g *Graph) String() string {
	nodes := ""
	for _, n := range g.nodes {
		nodes += fmt.Sprintf("%v ", *n)
	}
	return fmt.Sprintf("%s%v", nodes, g.edges)
}

// NewGraph constructs new Graph value
func NewGraph(N int) *Graph {
	nodes := make([]*Node, N)
	edges := make([][]int, N)

	for i := range nodes {
		nodes[i] = &Node{rand.Float64(), rand.Float64()}
	}
	for from := range edges {
		for to := from + 1; to < N; to++ {
			connected := rand.Float64()
			if connected < 0.5 {
				edges[from] = append(edges[from], to)
				edges[to] = append(edges[to], from)
			}
		}
	}

	return &Graph{nodes, edges}
}
