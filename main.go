package main

type GraphData struct {
	Location       []string
	DistanceMatrix [][]int
}

type ShortestResult struct {
	Path []int
	Cost int
}

func main() {
	//graphData := getMatrixDistanceFromGoogleMap(location)
	//writeGraphDataToFile(graphData)

	//graph := generateSampleData()
	//result := shortestPathUsingDijkstra(graph, 0, 5)
	//printShortestPath("Dijkstra", result, graph.Location)
	//
	//result = shortestPathUsingBellmanFord(graph, 0, 5)
	//printShortestPath("BellmanFord", result, graph.Location)
	//
	//result = shortestPathUsingFloyWarshal(graph, 0, 5)
	//printShortestPath("FloyWarshal", result, graph.Location)
	//
	//kthShortestPathUsingYen(graph, 3, 0, 5, shortestPathUsingDijkstra, "Djikstra")
	//kthShortestPathUsingYen(graph, 3, 0, 5, shortestPathUsingBellmanFord, "BellmanFord")
	//kthShortestPathUsingYen(graph, 3, 0, 5, shortestPathUsingFloyWarshal, "FloyWarshal")
	//
	graph := loadGraphDataFromFile()
	kthShortestPathUsingYen(graph, 3, 0, 1, shortestPathUsingFloyWarshal, "FloyWarshal")
}
