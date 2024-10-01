package graph

type Edges struct {
	fromVertex *Vertex
	toVertex   *Vertex
	weight     *int
}

func NewEdge(from, to *Vertex, w *int) *Edges {
	return &Edges{
		fromVertex: from,
		toVertex:   to,
		weight:     w,
	}
}

func (e *Edges) GetStartVertex() *Vertex {
	return e.fromVertex
}

func (e *Edges) GetEndVertex() *Vertex {
	return e.toVertex
}

func (e *Edges) GetWeight() *int {
	return e.weight
}

// Below are the utilities to deep check the equality of array of ages
type EdgesSlice []*Edges

func (edges EdgesSlice) isEqual(other EdgesSlice) bool {
	if len(edges) != len(other) {
		return false
	}
	for i, edge := range edges {
		if !edge.isEqual(other[i]) {
			return false
		}
	}
	return true
}

// Custom method to compare Edges structs by content
func (e *Edges) isEqual(other *Edges) bool {
	if e == nil || other == nil {
		return e == other
	}
	return e.fromVertex.isEqual(other.fromVertex) && e.toVertex.isEqual(other.toVertex) && e.weight == other.weight
}
