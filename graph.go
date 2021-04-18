package main

import (
	"container/list"
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

func MockEdges() (n int, edges []*Edge) {
	edges = []*Edge{
		{0, 1, 4},
		{0, 3, 8},
		{1, 4, 1},
		{1, 2, 2},
		{4, 2, 3},
		{2, 5, 3},
		{3, 4, 2},
	}
	for _, e := range edges {
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

func DFS(g *Graph, st int) {
	f := make([]bool, g.n)
	dFS(g, f, st, []int{-1}, -1)
	fmt.Println()
}

func FindConnectedComponents(g *Graph) (r int, cpn []int) {
	f := make([]bool, g.n)
	cpn = make([]int, g.n)
	for v := 0; v < g.n; v++ {
		if !f[v] {
			r++
			dFS(g, f, v, cpn, r)
		}
	}
	return r, cpn
}

func dFS(g *Graph, f []bool, v int, cpn []int, id int) {
	if f[v] {
		return
	}
	f[v] = true
	cpn[v] = id
	fmt.Print(v, "->")
	for _, node := range g.adjOf[v] {
		dFS(g, f, node.k, cpn, id)
	}
}

func BFS(g *Graph, st int) {
	f := make([]bool, g.n)
	q := list.New()
	bFS(g, f, q, st)
}

func bFS(g *Graph, f []bool, q *list.List, v int) {
	if f[v] {
		return
	}
	f[v] = true
	fmt.Println(v)

}
