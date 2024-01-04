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

func  (dll *doublyLinkedList) append(data int) {
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
    	ddl.append(data)
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

func (ddl* doublyLinkedList) remove(data int, idx int) (int, error) {
    if idx == 0 {
        ddl.head = ddl.head.next
    } else {
        currNode := ddl.head
        for i := 0; i < idx-1; i++ {
            currNode = currNode.next
        }
        currNode.next = currNode.next.next
    }
    return 0, nil
}

func (ddl* doublyLinkedList) prepend(data int) (int, error) {
    newNode := &node{
        data: data,
        prev: nil,
        next: nil,
    }
    newNode.next = ddl.head
    ddl.head.prev = newNode
    ddl.head = newNode
    return 0, nil
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
    dll.append(12)
    dll.append(26)
    dll.append(8)
    dll.append(78)

    fmt.Println("ddl forward:")
    dll.PrintForward()

    fmt.Println("ddl reverse:")
    dll.PrintReverse()
    fmt.Println("ddl insertAt:")
    dll.insertAt(100, 2)
    dll.PrintForward()
    fmt.Println("ddl remove:")
    dll.remove(100, 2)
    dll.PrintForward()
    dll.prepend(120)
    dll.PrintForward()
}
