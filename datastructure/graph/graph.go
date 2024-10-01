package graph

import (
	"fmt"
)

type Graph struct {
	vertices   []*Vertex
	isWeighted bool
	isDirected bool
}

func NewGraph(isWeighted, isDirected bool) *Graph {
	return &Graph{
		vertices:   []*Vertex{},
		isWeighted: isWeighted,
		isDirected: isDirected,
	}
}

func (g *Graph) AddVertex(d string) *Vertex {
	if Contains(g.vertices, d) {
		fmt.Printf("\nVertex: %s already present in graph", d)
		return nil
	}
	newVertex := NewVertex(d)
	g.vertices = append(g.vertices, newVertex)
	return newVertex
}

func (g *Graph) AddEdges(v1, v2 *Vertex, weight *int) {
	if !g.isWeighted {
		weight = nil
	}
	v1.addEdges(v2, weight)
	if !g.isDirected {
		v2.addEdges(v1, weight)
	}
}

func (g *Graph) RemoveEdge(v1, v2 *Vertex) {
	v1.removeEdges(v2)
	if !g.isDirected {
		v2.removeEdges(v1)
	}
}

func (g *Graph) RemoveVertex(v *Vertex) {
	if !Contains(g.vertices, v.data) {
		fmt.Printf("\nVertex: %+v not present in graph", v)
		return
	}
	var updatedVertices []*Vertex
	for _, val := range g.vertices {
		if val == v {
			continue
		}
		updatedVertices = append(updatedVertices, val)
		//
		val.removeEdges(v)
	}
	g.vertices = updatedVertices
}

func (g *Graph) PrintGraph(showWeight bool) {
	for _, v := range g.vertices {
		v.print(showWeight)
	}
}

func (g *Graph) GetVertices() []*Vertex {
	return g.vertices
}

func (g *Graph) IsWeighted() bool {
	return g.isWeighted
}

func (g *Graph) IsDirected() bool {
	return g.isDirected
}

func Contains(s []*Vertex, val string) bool {
	for _, v := range s {
		if v.data == val {
			return true
		}
	}
	return false
}

func Weight(x any) *int {
	l, ok := x.(int)
	if !ok {
		return nil
	}
	return &l
}

func InitGraph() {
	gr := NewGraph(true, true)
	A := gr.AddVertex("SANTA ROSA")
	B := gr.AddVertex("SR Transit Mall")
	C := gr.AddVertex("Rohnert park")
	D := gr.AddVertex("COTATI HUB")
	E := gr.AddVertex("Petaluma")
	F := gr.AddVertex("Novato")
	G := gr.AddVertex("San Rafael")
	H := gr.AddVertex("Leave SanRafael")
	gr.AddEdges(A, B, Weight(20))
	gr.AddEdges(A, D, Weight(80))
	gr.AddEdges(B, F, Weight(10))
	gr.AddEdges(F, D, Weight(40))
	gr.AddEdges(D, G, Weight(20))
	gr.AddEdges(G, E, Weight(30))
	gr.AddEdges(E, B, Weight(50))
	gr.AddEdges(D, C, Weight(10))
	gr.AddEdges(C, F, Weight(50))
	gr.AddEdges(C, H, Weight(20))
	gr.PrintGraph(true)
	gr.RemoveVertex(D)
	fmt.Printf("\n\nAfter removing COTATI HUB")
	gr.PrintGraph(true)
	gr.RemoveEdge(E, B)
	fmt.Printf("\n\nAfter removing edge from Petaluma to SR Transit Mall")
	gr.PrintGraph(true)

	fmt.Println("***********		***********		***********")
	fmt.Println("Graph2:")
	gr2 := NewGraph(true, true)
	A1 := gr2.AddVertex("A")
	B1 := gr2.AddVertex("B")
	C1 := gr2.AddVertex("C")
	D1 := gr2.AddVertex("D")
	E1 := gr2.AddVertex("E")
	gr2.AddEdges(A1, B1, Weight(4))
	gr2.AddEdges(A1, C1, Weight(11))
	gr2.AddEdges(B1, C1, Weight(1))
	gr2.AddEdges(B1, D1, Weight(2))
	gr2.AddEdges(B1, E1, Weight(3))
	gr2.AddEdges(C1, E1, Weight(1))
	gr2.AddEdges(E1, D1, Weight(2))
	gr2.AddEdges(D1, B1, Weight(3))
	gr2.PrintGraph(true)

	path, distance := GetShortestPath(A1, E1, gr2)

	fmt.Printf("Shortest distance from %s to %s: %d \nFollowed Path: %s", A1.data, E1.data, distance, path)
}
