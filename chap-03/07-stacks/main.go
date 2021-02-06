package main

import "fmt"

type Element struct {
	elementValue int
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func New() *Stack {
	stack := &Stack{
		elements:     make([]*Element, 0),
		elementCount: 0,
	}

	return stack
}

func (s *Stack) Push(element *Element) {
	s.elements = append(s.elements[:s.elementCount], element)
	s.elementCount = s.elementCount + 1
}

func (s *Stack) Pop() *Element {
	if s.elementCount == 0 {
		return nil
	}

	lastItem := s.elements[s.elementCount-1]
	s.elements = s.elements[:s.elementCount-1]
	s.elementCount = s.elementCount - 1
	return lastItem
}

func main() {
	stack := New()

	stack.Push(&Element{1})
	stack.Push(&Element{2})
	stack.Push(&Element{3})
	stack.Push(&Element{4})

	fmt.Println("Items in stack:", stack.elementCount)
	for _, val := range stack.elements {
		fmt.Println(*val)
	}

	popedItem := stack.Pop()
	fmt.Println("Poped item 1:", *popedItem)

	fmt.Println("Items in stack after popping:", stack.elementCount)
	for _, val := range stack.elements {
		fmt.Println(*val)
	}

	popedItem = stack.Pop()
	fmt.Println("Poped item 2:", *popedItem)

	fmt.Println("Items in stack after popping:", stack.elementCount)
	for _, val := range stack.elements {
		fmt.Println(*val)
	}
}
