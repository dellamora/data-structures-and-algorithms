package main

import "fmt"
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("empty queue")
	} 

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) Size() int {
	return len(q.items)
}

// func main() {
// 	myQueue := Queue{}

// 	myQueue.Enqueue(1)
// 	myQueue.Enqueue(2)
// 	myQueue.Enqueue(3)
// 	myQueue.Enqueue(4)
// 	myQueue.Enqueue(5)

// 	fmt.Println("items:", myQueue.items)

// 	myQueue.Dequeue()
// 	myQueue.Dequeue()
// 	myQueue.Dequeue()
// 	fmt.Println(myQueue.items)
	
// }