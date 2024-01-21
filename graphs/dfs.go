package main

// on a djacent list

type GraphDFS struct {
	adjacencyList map[int][]int
}

func walk(graph *GraphDFS, curr int, needle int, seen map[int]bool) []int {
	if curr == needle {
		return []int{curr}
	}
	if seen[curr] {
		return nil
	}
	seen[curr] = true

	//recurse
	var list []int = graph.adjacencyList[curr]
	for _, neighbor := range list {
		path := walk(graph, neighbor, needle, seen)
		if path != nil {
			return append([]int{curr}, path...)
		}
	}

	return nil
}

func dfs(graph *GraphDFS, source int, needle int) []int {
	seen := make(map[int]bool)
	return walk(graph, source, needle, seen)
}

// func main() {
// 	graph := &Graph{
// 		adjacencyList: map[int][]int{
// 			0: []int{1, 2},
// 			1: []int{0, 3},
// 			2: []int{0, 3, 4},
// 			3: []int{1, 2, 4},
// 			4: []int{2, 3},
// 		},
// 	}

// 	path := dfs(graph, 0, 3)
// 	if len(path) > 0 {
// 		fmt.Println("o/ path found:", path)
// 	} else {
// 		fmt.Println("path not found lol.")
// 	}
// }
