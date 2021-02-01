package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iheap *IntegerHeap) Len() int {
	return len(*iheap)
}

func (iheap *IntegerHeap) Less(i, j int) bool {
	return (*iheap)[i] < (*iheap)[j]
}

func (iheap *IntegerHeap) Swap(i, j int) {
	(*iheap)[i], (*iheap)[j] = (*iheap)[j], (*iheap)[i]
}

func (iheap *IntegerHeap) Push(item interface{}) {
	*iheap = append(*iheap, item.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	// var previous IntegerHeap = *iheap
	length := len(*iheap)
	lastItem := (*iheap)[length-1]
	*iheap = (*iheap)[0 : length-1]
	return lastItem
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5, 6}

	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	heap.Push(intHeap, 7)

	fmt.Println("Minimum:", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("%d\n", intHeap.Pop())
	}
}
