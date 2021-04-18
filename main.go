package main

func main() {
	var g Graph
	n, edges := MockEdges()
	g.init(n, edges, true)
	g.print()
	DFS(&g, 4)
}
