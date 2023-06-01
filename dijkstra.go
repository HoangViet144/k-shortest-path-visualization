package main

import (
	"container/heap"
	"math"
)

func shortestPathUsingDijkstra(graph *GraphData, source, sink int) ShortestResult {
	V := len(graph.DistanceMatrix)

	parent := make([]int, V)
	for i := 0; i < V; i++ {
		parent[i] = -1
	}

	dist := make([]int, V)
	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt
	}

	adjList := convertAdjacencyMatrixToAdjacencyList(V, graph.DistanceMatrix)

	q := make(PriorityQueue, 0)
	heap.Init(&q)

	heap.Push(
		&q, &Item{
			value:    source,
			priority: 0,
		},
	)
	dist[source] = 0
	parent[source] = source

	for q.Len() > 0 {
		item := heap.Pop(&q).(*Item)
		u := item.value

		for _, vInfo := range adjList[u] {
			v := vInfo[0]
			weight := vInfo[1]
			if dist[v] > dist[u]+weight {
				dist[v] = dist[u] + weight
				parent[v] = u
				heap.Push(
					&q, &Item{
						value:    v,
						priority: dist[v],
					},
				)
			}
		}
	}

	if dist[sink] == math.MaxInt {
		return ShortestResult{
			Path: []int{},
			Cost: dist[sink],
		}
	}

	cur := sink
	shortestPath := make([]int, 0)
	for {
		shortestPath = append(shortestPath, cur)
		if parent[cur] == cur {
			break
		}
		cur = parent[cur]
	}

	for i := 0; i < len(shortestPath)/2; i++ {
		shortestPath[i], shortestPath[len(shortestPath)-1-i] = shortestPath[len(shortestPath)-1-i], shortestPath[i]
	}

	return ShortestResult{
		Path: shortestPath,
		Cost: dist[sink],
	}
}
