package main

import "fmt"

type Stack struct {
	items []int
}

func (q *Stack) push(item int) {
	q.items = append(q.items, item)
}


func (q *Stack) pop() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("Empty Stack")
	}
	item := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return item, nil
}


func (q *Stack) Size() int {
	return len(q.items)
}


func main() {
	myStack := Stack{}

	myStack.push(1)
	myStack.push(2)
	myStack.push(3)
	myStack.push(4)
	myStack.push(5)

	fmt.Println("items:", myStack.items)

	myStack.pop()
	fmt.Println("items:", myStack.items)
}