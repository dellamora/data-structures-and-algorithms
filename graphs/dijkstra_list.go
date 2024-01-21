package main

import (
	"fmt"
	"math"
)

type WeightedAdjacencyList struct {
	adjacencyList map[int]map[int]int
}

func minDist(seen []bool, dists []int) int {
	min := math.MaxInt
	minIndex := -1

	for i, val := range dists {
		if !seen[i] && val < min {
			min = val
			minIndex = i
		}
	}

	return minIndex
}

func hasUnvisited(seen []bool) bool {
	for _, val := range seen {
		if !val {
			return true
		}
	}
	return false
}

func dijkstraList(source int, sink int, arr *WeightedAdjacencyList) []int {
	seen := make([]bool, len(arr.adjacencyList))
	dists := make([]int, len(arr.adjacencyList))
	for i := range dists {
		dists[i] = math.MaxInt
	}
	dists[source] = 0

	for hasUnvisited(seen) {
		curr := minDist(seen, dists)
		seen[curr] = true

		for neighbor, weight := range arr.adjacencyList[curr] {
			if !seen[neighbor] {
				newDist := dists[curr] + weight
				if newDist < dists[neighbor] {
					dists[neighbor] = newDist
				}
			}
		}
	}

	return dists
}

func main() {

	adjList := WeightedAdjacencyList{
		adjacencyList: map[int]map[int]int{
			0: {1: 4, 2: 1},
			1: {0: 4, 2: 2, 3: 5},
			2: {0: 1, 1: 2, 3: 1},
			3: {1: 5, 2: 1},
		},
	}

	source := 0
	sink := 3
	result := dijkstraList(source, sink, &adjList)

	fmt.Printf("Shortest distances from node %d to all other nodes: %v\n", source, result)
}
