package graph

import (
	"container/heap"
	"math"
)

func GetShortestPath(startVertex, endVertex *Vertex, g *Graph) (string, int) {
	if startVertex == endVertex {
		return startVertex.data, 0
	}
	distance := make(map[string]int)
	prevsVertex := make(map[string]*Vertex)
	priorityQ := make(PriorityQueue, 0)
	heap.Init(&priorityQ)
	heap.Push(&priorityQ, NewGraphQueue(startVertex, 0))
	distance[startVertex.data] = 0
	prevsVertex[startVertex.data] = NewVertex("Nil")
	visitedMap := make(map[string]struct{})

	for _, vertex := range g.vertices {
		if vertex != startVertex {
			distance[vertex.data] = math.MaxInt32
		}
	}

	for !priorityQ.isEmpty() {
		currentVertex := heap.Pop(&priorityQ).(*GraphPriorityQueue).vertex

		if _, ok := visitedMap[currentVertex.data]; ok {
			continue
		}

		visitedMap[currentVertex.data] = struct{}{}

		for _, edges := range currentVertex.edges {
			newDistanceFromCurrentVertex := *edges.weight + distance[currentVertex.data]
			adjacentVertex := edges.toVertex.data
			if newDistanceFromCurrentVertex < distance[adjacentVertex] {
				distance[adjacentVertex] = newDistanceFromCurrentVertex
				prevsVertex[adjacentVertex] = currentVertex
				heap.Push(&priorityQ, NewGraphQueue(edges.toVertex, distance[adjacentVertex]))
			}
		}
	}

	var path string
	path = endVertex.data

	pathVertex := prevsVertex[endVertex.data]

	for pathVertex.data != "Nil" {
		path = pathVertex.data + " --> " + path
		pathVertex = prevsVertex[pathVertex.data]
	}
	return path, distance[endVertex.data]
}
