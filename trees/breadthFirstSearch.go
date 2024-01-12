package main

import "fmt"

// type node struct {
// 	data  int
// 	left  *node
// 	right *node
// }


type Queue struct {
	items []*node
}

func (q *Queue) Enqueue(item *node) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (*node, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("empty queue")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) Size() int {
	return len(q.items)
}

func bfs(head *node, needle int) bool {
	var q = Queue{}
	q.Enqueue(head)

	for q.Size() > 0 {
		curr, err := q.Dequeue()
		if err != nil {
			fmt.Println(err)
			return false
		}

		if curr != nil {
			if curr.data == needle {
				return true
			}

			if curr.left != nil {
				q.Enqueue(curr.left)
			}

			if curr.right != nil {
				q.Enqueue(curr.right)
			}
		}
	}
	return false
}


// func main() {
// 	root := &node{
// 		data: 1,
// 		left: &node{
// 			data: 2,
// 			left: &node{
// 				data: 4,
// 				left: nil,
// 				right: nil,
// 			},
// 			right: &node{
// 				data: 5,
// 				left: nil,
// 				right: nil,
// 			},
// 		},
// 		right: &node{
// 			data: 3,
// 			left: &node{
// 				data: 6,
// 				left: nil,
// 				right: nil,
// 			},
// 			right: &node{
// 				data: 7,
// 				left: nil,
// 				right: nil,
// 			},
// 		},
// 	}

// 	fmt.Println(bfs(root, 6))
// }
