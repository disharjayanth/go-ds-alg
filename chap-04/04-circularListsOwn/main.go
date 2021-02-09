package main

import "fmt"

type CircularLinkedList struct {
	head *Node
}

func (cl *CircularLinkedList) InsertNewNode(val int) {
	if cl.head == nil {
		cl.head = &Node{
			val:  val,
			next: nil,
		}
		cl.head.next = cl.head
		return
	}

	for node := cl.head; node != nil; node = node.next {
		if node.next == cl.head {
			node.next = &Node{
				val:  val,
				next: cl.head,
			}
			return
		}
	}
}

func (cl *CircularLinkedList) print() {
	for node := cl.head; node != nil; node = node.next {
		fmt.Println("Node", node, "its next node:", node.next)
		if node.next == cl.head {
			return
		}
	}
}

type Node struct {
	val  int
	next *Node
}

func main() {
	cl := &CircularLinkedList{}

	cl.InsertNewNode(2)
	cl.InsertNewNode(3)
	cl.InsertNewNode(4)
	cl.InsertNewNode(5)
	cl.InsertNewNode(6)
	cl.InsertNewNode(7)

	cl.print()
}
