package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"
)

type FindShortestPathFunc func(graph *GraphData, source, sink int) ShortestResult

// Ref: https://en.wikipedia.org/wiki/Yen%27s_algorithm, https://www.linchenguang.com/2018/01/30/Yen-s-algorithm
func kthShortestPathUsingYen(
	data *GraphData,
	K int,
	source,
	sink int,
	findShortestPathFunc FindShortestPathFunc,
	shortestAlgorithm string,
) []ShortestResult {
	startTime := time.Now()

	V := len(data.DistanceMatrix)
	A := make([]ShortestResult, 0)

	// Determine the shortest path from the source to the sink.
	result := shortestPathUsingBellmanFord(copyGraphData(data), source, sink)
	A = append(A, result)

	// Initialize the set to store the potential kth shortest path.
	B := make([]ShortestResult, 0)
	ExistInB := make(map[string]struct{})

	for k := 1; k < K; k++ {
		cpData := copyGraphData(data)

		// The spur node ranges from the first node to the next to last node in the previous k-shortest path.
		for i := 0; i < len(A[k-1].Path)-1; i++ {
			// Spur node is retrieved from the previous k-shortest path, k − 1.
			spurNode := A[k-1].Path[i]
			// The sequence of nodes from the source to the spur node of the previous k-shortest path.
			rootPath := make([]int, i+1)
			copy(rootPath, A[k-1].Path[0:i+1])

			// Remove the links that are part of the previous shortest paths which share the same root path
			for _, p := range A {
				if reflect.DeepEqual(rootPath, p.Path[0:i+1]) {
					cpData.DistanceMatrix[p.Path[i]][p.Path[i+1]] = math.MaxInt
				}
			}

			// Remove node in rootPath except spurNode
			for u := 0; u < V; u++ {
				for _, v := range rootPath {
					if v == spurNode {
						continue
					}
					cpData.DistanceMatrix[u][v] = math.MaxInt
				}
			}

			for _, v := range rootPath {
				if v == spurNode {
					continue
				}
				for u := 0; u < V; u++ {
					cpData.DistanceMatrix[v][u] = math.MaxInt
				}
			}

			// Calculate the spur path from the spur node to the sink.
			// Consider also checking if any spurPath found
			spurResult := findShortestPathFunc(cpData, spurNode, sink)
			if spurResult.Cost == math.MaxInt {
				continue
			}

			// Entire path is made up of the root path and spur path.
			totalPath := append(rootPath[0:len(rootPath)-1], spurResult.Path...)

			cost := 0
			for u := 0; u < len(totalPath)-1; u++ {
				cost += data.DistanceMatrix[totalPath[u]][totalPath[u+1]]
			}

			// Add the potential k-shortest path to the heap.
			potentialPathKey, _ := json.Marshal(totalPath)
			if _, exist := ExistInB[string(potentialPathKey)]; !exist {
				B = append(
					B, ShortestResult{
						Path: totalPath,
						Cost: cost,
					},
				)
				ExistInB[string(potentialPathKey)] = struct{}{}
			}

			// Add back the edges and nodes that were removed from the graph.
			cpData = copyGraphData(data)
		}

		if len(B) == 0 {
			// This handles the case of there being no spur paths, or no spur paths left.
			// This could happen if the spur paths have already been exhausted (added to A),
			// or there are no spur paths at all - such as when both the source and sink vertices
			// lie along a "dead end".
			break
		}

		// Sort the potential k-shortest paths by cost.
		sort.Slice(
			B, func(m, n int) bool {
				return B[m].Cost < B[n].Cost
			},
		)

		// Add the lowest cost path becomes the k-shortest path.
		A = append(A, copyShortestResult(B[0]))

		// In fact we should rather use shift since we are removing the first element
		B = B[1:]
	}

	endTime := time.Now()

	for i, kthShortestResult := range A {
		fmt.Printf("The %d-shorthest path using Yen and %s algorithm:\n", i+1, shortestAlgorithm)
		for j := 0; j < len(kthShortestResult.Path); j++ {
			fmt.Printf("[%s]", data.Location[kthShortestResult.Path[j]])
			if j < len(kthShortestResult.Path)-1 {
				fmt.Printf(" -> ")
			} else {
				fmt.Println()
			}
		}
		fmt.Println("Cost:", kthShortestResult.Cost)
	}
	fmt.Println("Execution time:", endTime.Sub(startTime))
	fmt.Println("---")

	return A
}
