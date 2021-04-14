package main

func main() {
	var g Graph
	n, edges := ReadEdges()
	g.init(n, edges, true)
	g.print()
}
