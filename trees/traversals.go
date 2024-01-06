package main

import "fmt"

type node struct { 
	data int
	left *node
	right *node
}

func visitNode(node *node) {
	fmt.Println(node.data)
}

func preOrder(node *node) {
	if node == nil {
		return
	}
	visitNode(node)
	preOrder(node.left)
	preOrder(node.right)
}

func inOrder(node *node) {
	if node == nil {
		return
	}
	inOrder(node.left)
	visitNode(node)
	inOrder(node.right)
}

func postOrder(node *node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	visitNode(node)
}

func main() {
	root := &node{
		data: 1,
		left: &node{
			data: 2,
			left: &node{
				data: 4,
				left: nil,
				right: nil,
			},
			right: &node{
				data: 5,
				left: nil,
				right: nil,
			},
		},
		right: &node{
			data: 3,
			left: &node{
				data: 6,
				left: nil,
				right: nil,
			},
			right: &node{
				data: 7,
				left: nil,
				right: nil,
			},
		},
	}
	fmt.Println("preOrder")
	preOrder(root)
	fmt.Println("inOrder")
	inOrder(root)
	fmt.Println("postOrder")
	postOrder(root)
}