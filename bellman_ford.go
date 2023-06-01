package main

import "math"

func shortestPathUsingBellmanFord(graph *GraphData, source, sink int) ShortestResult {
	V := len(graph.DistanceMatrix)

	dist := make([]int, V)
	parent := make([]int, V)
	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt
		parent[i] = -1
	}

	dist[source] = 0
	parent[source] = source

	adjList := convertAdjacencyMatrixToAdjacencyList(V, graph.DistanceMatrix)

	for i := 0; i < V; i++ {
		for u := 0; u < V; u++ {
			for _, item := range adjList[u] {
				v := item[0]
				weight := item[1]
				if weight != math.MaxInt && dist[u] != math.MaxInt && dist[u]+weight < dist[v] {
					dist[v] = dist[u] + weight
					parent[v] = u
				}
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
