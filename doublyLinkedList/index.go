package main

import "fmt"

type node struct {
	data int
	node string
	prev *node
	next *node
}


type doublyLinkedList struct {
	head *node
	tail *node
}

func  (dll *doublyLinkedList) AddNode(data int) {
    newNode := &node{
        data: data,
        prev: nil,
        next: nil,
    }
    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
    } else {
        newNode.prev = dll.tail
        dll.tail.next = newNode
        dll.tail = newNode
    }
}


func (dll *doublyLinkedList) PrintForward() {
    currNode := dll.head
    for currNode != nil {
        fmt.Printf("%d ", currNode.data)
        currNode = currNode.next
    }
    fmt.Println()
}

func (dll *doublyLinkedList) PrintReverse() {
    currNode := dll.tail
    for currNode != nil {
        fmt.Printf("%d ", currNode.data)
        currNode = currNode.prev
    }
    fmt.Println()
}


func main() {
    dll := doublyLinkedList{}
    dll.AddNode(12)
    dll.AddNode(26)
    dll.AddNode(8)
    dll.AddNode(78)

    fmt.Println("ddl forward:")
    dll.PrintForward()

    fmt.Println("ddl reverse:")
    dll.PrintReverse()
}
