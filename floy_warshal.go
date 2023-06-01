package main

import "math"

func shortestPathUsingFloyWarshal(graph *GraphData, source, sink int) ShortestResult {
	cpGraph := copyGraphData(graph)

	V := len(cpGraph.DistanceMatrix)

	next := make([][]int, V)
	for i := 0; i < V; i++ {
		next[i] = make([]int, V)
	}

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if cpGraph.DistanceMatrix[i][j] != math.MaxInt {
				next[i][j] = j
			} else {
				next[i][j] = -1
			}
		}
	}

	for k := 0; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				if cpGraph.DistanceMatrix[i][k] == math.MaxInt || cpGraph.DistanceMatrix[k][j] == math.MaxInt {
					continue
				}
				if cpGraph.DistanceMatrix[i][j] > cpGraph.DistanceMatrix[i][k]+cpGraph.DistanceMatrix[k][j] {
					cpGraph.DistanceMatrix[i][j] = cpGraph.DistanceMatrix[i][k] + cpGraph.DistanceMatrix[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}

	if cpGraph.DistanceMatrix[source][sink] == math.MaxInt {
		return ShortestResult{
			Path: []int{},
			Cost: math.MaxInt,
		}
	}

	cur := source
	shortestPath := make([]int, 0)
	shortestPath = append(shortestPath, cur)
	for cur != sink {
		cur = next[cur][sink]
		shortestPath = append(shortestPath, cur)
	}

	return ShortestResult{
		Path: shortestPath,
		Cost: cpGraph.DistanceMatrix[source][sink],
	}
}
