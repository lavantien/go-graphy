package main

import (
	"fmt"
)

type Edge struct {
	a int
	b int
	w float64
}

type Node struct {
	k   int
	val float64
}

type Graph struct {
	n     int
	adjOf [][]*Node
}

func NewGraph(n int) (g Graph) {
	g.n = n
	g.adjOf = make([][]*Node, n)
	return g
}

func ReadEdges() (n int, edges []*Edge) {
	var t int
	fmt.Scanln(&t)
	edges = make([]*Edge, t)
	for i := range edges {
		e := Edge{}
		fmt.Scanln(&e.a, &e.b, &e.w)
		edges[i] = &e
		if e.a > n {
			n = e.a
		}
		if e.b > n {
			n = e.b
		}
	}
	return n + 1, edges
}

func (g *Graph) init(n int, edges []*Edge, bi bool) {
	*g = NewGraph(n)
	for _, e := range edges {
		node := Node{e.b, e.w}
		g.adjOf[e.a] = append(g.adjOf[e.a], &node)
		if bi {
			revNode := Node{e.a, e.w}
			g.adjOf[e.b] = append(g.adjOf[e.b], &revNode)
		}
	}
}

func (g *Graph) print() {
	for i, head := range g.adjOf {
		fmt.Print(i, ":")
		for _, node := range head {
			fmt.Print("(", node.k, ",", node.val, ")")
		}
		fmt.Println()
	}
}
