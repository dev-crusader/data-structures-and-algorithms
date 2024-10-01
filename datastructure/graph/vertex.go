package graph

import (
	"fmt"
	"strconv"
)

type Vertex struct {
	data  string
	edges []*Edges
}

func NewVertex(val string) *Vertex {
	return &Vertex{
		data:  val,
		edges: []*Edges{},
	}
}

func (v *Vertex) addEdges(inputVertex *Vertex, weight *int) {
	v.edges = append(v.edges, NewEdge(v, inputVertex, weight))
}

func (v *Vertex) removeEdges(vertexRemove *Vertex) {
	var updatedEdges []*Edges
	for _, vtx := range v.edges {
		if vtx.toVertex == vertexRemove {
			continue
		}
		updatedEdges = append(updatedEdges, vtx)
	}
	v.edges = updatedEdges
}

func (v *Vertex) GetData() string {
	return v.data
}

func (v *Vertex) GetEdges() []*Edges {
	return v.edges
}

func (v *Vertex) print(showWeight bool) {
	fmt.Println("\nVertex: ", v.data)
	for _, edge := range v.edges {
		result := " --> " + edge.toVertex.data
		if showWeight {
			var w string
			if edge.weight == nil {
				w = "Nil"
			} else {
				w = strconv.Itoa(*edge.weight)
			}
			result += fmt.Sprintf(" (%s)", w)
		}

		fmt.Println(result)
	}
	fmt.Println()
}

// compare two vertex

func (v *Vertex) isEqual(otherVertex *Vertex) bool {
	if v == nil || otherVertex == nil {
		return v == otherVertex
	}

	t := EdgesSlice(v.edges)
	return v.data == otherVertex.data && t.isEqual(otherVertex.edges)
}
