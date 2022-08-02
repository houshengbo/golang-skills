package graph

import "fmt"

var (
	Player string = "玩家"
)

// A graph organizes items in an interconnected network.
// Each item is a node (or vertex). Nodes are connected by edges.
// A graph is composed of nodes (or vertices) that are connected by edges.

// Strengths:
//
// Representing links. Graphs are ideal for cases where you're working with things that connect to other things. Nodes
// and edges could, for example, respectively represent cities and highways, routers and ethernet cables, or Facebook
// users and their friendships.
//
// Weaknesses:
//
// Scaling challenges. Most graph algorithms are O(n∗lg(n))O(n*lg(n))O(n∗lg(n)) or even slower. Depending on the size
// of your graph, running algorithms across your nodes may not be feasible.
// This is an implementation of an adjacency list.

// Adjacency list VS adjacency matrix
//

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) AddVertex(k string) {
	if !contains(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(from, to string) {
	// Get the vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// Check the error
	if fromVertex == nil || toVertex == nil {
		return
	} else if !contains(fromVertex.adjacency, toVertex.key) {
		fromVertex.adjacency = append(fromVertex.adjacency, toVertex)
	}
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n %s %v : ", Player, v.key)
		for _, v := range v.adjacency {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}

func (g *Graph) getVertex(k string) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, key string) bool {
	for _, v := range s {
		if key == v.key {
			return true
		}
	}
	return false
}

type Vertex struct {
	key       string
	adjacency []*Vertex
}
