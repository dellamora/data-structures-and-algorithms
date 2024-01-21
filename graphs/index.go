package main

import "fmt"

// Graph Terms

// cycle: When you start at Node(x), follow the links, and back at Node(x) again.
// acyclic: A graph that contains no cycles
// connected: When every node has a path to another node
// Directed: When the edges have a direction
// Undirected: When the edges have no direction
// Weighted: When the edges have a weight
// dag: Directed Acyclic Graph

// Implementation Terms

// node: A single item in a graph
// edge: A connection between two nodes


type Graph struct {
	adjacencyMatrix [][]int
}

func bfs(graph *Graph, startNode int) []int {
	numNodes := len(graph.adjacencyMatrix)
	visited := make([]bool, numNodes)
	queue := []int{startNode}
	result := []int{startNode}

	visited[startNode] = true

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for neighbor, isEdge := range graph.adjacencyMatrix[currentNode] {
			if isEdge == 1 && !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				result = append(result, neighbor)
			}
		}
	}

	return result
}

func main() {
	graph := &Graph{
		adjacencyMatrix: [][]int{
			{0, 1, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{1, 0, 0, 1, 1},
			{0, 1, 1, 0, 1},
			{0, 0, 1, 1, 0},
		},
	}

	startNode := 2
	result := bfs(graph, startNode)

	fmt.Println(result)
}