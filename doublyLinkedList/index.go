package main

import (
	"fmt"
)

type node struct {
	data int
	prev *node
	next *node
}


type doublyLinkedList struct {
	head *node
	tail *node
}

func  (dll *doublyLinkedList) prepend(data int) {
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

func (ddl* doublyLinkedList) insertAt(data int, idx int) (int, error) {
	newNode := &node{
        data: data,
        prev: nil,
        next: nil,
    }  
    if idx == 0 {
    	ddl.prepend(data)
    } else {
    	currNode := ddl.head
    	for i := 0; i < idx-1; i++ {
    		currNode = currNode.next
    	}
    	newNode.next = currNode.next
    	currNode.next = newNode
    }
    return 0, nil
}

func (ddl* doublyLinkedList) remove() {

}

func (ddl* doublyLinkedList) append() {

}

func (ddl* doublyLinkedList) get() {
		
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
    dll.prepend(12)
    dll.prepend(26)
    dll.prepend(8)
    dll.prepend(78)

    fmt.Println("ddl forward:")
    dll.PrintForward()

    fmt.Println("ddl reverse:")
    dll.PrintReverse()
    dll.insertAt(100, 2)
    dll.PrintForward()
}
