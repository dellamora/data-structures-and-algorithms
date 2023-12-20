package main

import "fmt"

type Point struct {
	x, y int
}

type Stack struct {
	items []Point
}

func (q *Stack) push(item Point) {
	q.items = append(q.items, item)
}

func (q *Stack) pop() (Point, error) {
	if len(q.items) == 0 {
		return Point{}, fmt.Errorf("Empty Stack")
	}
	item := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return item, nil
}

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze [][]string, wall string, curr Point, end Point, seen [][]bool, path *Stack) bool {
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >=len(maze) {
		return false
	}

	if maze[curr.y][curr.x] == wall {
		return false
	}

	if curr.x == end.x && curr.y == end.y {
		return true
	}

	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	path.push(curr)
	for i := 0; i < len(directions); i++ {
		x,y := directions[i][0], directions[i][1]
		newPoint := Point{curr.x + x, curr.y + y} 
		if walk(maze, wall, newPoint, end, seen, path) {
			return true
		}
	}
	path.pop()
	return false
}

func Solve(maze [][]string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	var path Stack 

	walk(maze, wall, start, end, seen, &path)

	return path.items
}

func main() {
	start := Point{1,1}  
	end := Point{4, 1}
	maze := [][]string{
		{"#", "#", "#", "#", " "},
		{"#", " ", " ", " ", " "},
		{"#", " ", "#", "#", "#"},
		{"#", " ", " ", " ", "#"},
		{"#", " ", "#", "#", "#"},
	}

	result := Solve(maze, "#", start, end)
	fmt.Println(result)
	fmt.Println("Maze:")
	for _, row := range maze {
		fmt.Println(row)
	}

}