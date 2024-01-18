package main

import (
	"fmt"
)

type MinHeap struct {
	arr []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Insert(value int) {
	h.arr = append(h.arr, value)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.arr) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	min := h.arr[0]
	lastIndex := len(h.arr) - 1

	h.arr[0] = h.arr[lastIndex]
	h.arr = h.arr[:lastIndex]

	if len(h.arr) > 1 {
		h.heapifyDown(0)
	}

	return min, nil
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.arr[index] < h.arr[parentIndex] {
			h.arr[index], h.arr[parentIndex] = h.arr[parentIndex], h.arr[index]
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.arr) - 1

	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2

		smallest := index

		if leftChildIndex <= lastIndex && h.arr[leftChildIndex] < h.arr[smallest] {
			smallest = leftChildIndex
		}

		if rightChildIndex <= lastIndex && h.arr[rightChildIndex] < h.arr[smallest] {
			smallest = rightChildIndex
		}

		if smallest != index {
			h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index]
			index = smallest
		} else {
			break
		}
	}
}

func main() {
	heap := NewMinHeap()

	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(88)
	heap.Insert(2)

	fmt.Println(heap.arr)

	min, err := heap.ExtractMin()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(min)
		fmt.Println(heap.arr)
	}
}
