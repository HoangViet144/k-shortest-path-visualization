package main

import (
	"fmt"
	"math"
)

func convertAdjacencyMatrixToAdjacencyList(V int, adjacencyMatrix [][]int) [][][]int {
	adjList := make([][][]int, V)
	for i := 0; i < V; i++ {
		adjList[i] = make([][]int, 0)
	}

	for u := 0; u < V; u++ {
		for v := 0; v < V; v++ {
			if u == v {
				continue
			}
			weight := adjacencyMatrix[u][v]
			if weight != math.MaxInt {
				adjList[u] = append(adjList[u], []int{v, weight})
			}
		}
	}

	return adjList
}

func printShortestPath(algorithm string, shortestResult ShortestResult, location []string) {
	if shortestResult.Cost == math.MaxInt {
		fmt.Printf("There is no shortest path using: %s\n", algorithm)
	} else {
		fmt.Printf("Shortest path using: %s\n", algorithm)
		for i := 0; i < len(shortestResult.Path); i++ {
			fmt.Printf("[%s]", location[shortestResult.Path[i]])
			if i < len(shortestResult.Path)-1 {
				fmt.Printf(" -> ")
			} else {
				fmt.Println()
			}
		}
		fmt.Printf("Cost: %d\n", shortestResult.Cost)
	}

	fmt.Println("---")
}

func copyGraphData(graph *GraphData) *GraphData {
	V := len(graph.DistanceMatrix)
	cpData := &GraphData{
		Location:       graph.Location,
		DistanceMatrix: make([][]int, V),
	}

	for i := 0; i < V; i++ {
		cpData.DistanceMatrix[i] = make([]int, V)
	}

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			cpData.DistanceMatrix[i][j] = graph.DistanceMatrix[i][j]
		}
	}

	return cpData
}

func copyShortestResult(result ShortestResult) ShortestResult {
	cpResult := ShortestResult{
		Cost: result.Cost,
		Path: make([]int, len(result.Path)),
	}
	copy(cpResult.Path, result.Path)
	return cpResult
}
