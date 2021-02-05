package main

import "fmt"

type Node struct {
	property int
	prevNode *Node
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.property = property
	node.prevNode = nil
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.prevNode = node
	}

	linkedList.headNode = node
}

func (linkedList *LinkedList) IterateThroughList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(*node)
	}
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) LastNode() *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) AddToEnd(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil
	node.prevNode = nil

	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.prevNode = lastNode
	}
}

func (linkedList *LinkedList) AddAfterThisNode(nodeProperty int, property int) {
	node := &Node{}
	node.property = property
	node.prevNode = nil
	node.nextNode = nil

	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.prevNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) NodeBetweenTwoNodes(firstProperty int, secondProperty int) *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.prevNode != nil && node.nextNode != nil {
			if node.prevNode.property == firstProperty && node.nextNode.property == secondProperty {
				return node
			}
		}
	}
	return nil
}

func main() {
	doublyLinkedList := LinkedList{}

	fmt.Println("New headNode: 1")
	doublyLinkedList.AddToHead(1)
	doublyLinkedList.AddToHead(2)
	doublyLinkedList.AddToHead(3)
	doublyLinkedList.AddToHead(4)

	doublyLinkedList.IterateThroughList()

	fmt.Println("Node with property value 2 :", doublyLinkedList.NodeWithValue(2))

	fmt.Println("Last node in doubly linked list:", doublyLinkedList.LastNode())

	fmt.Println("Adding node 5 to end of doubly linked list:")
	doublyLinkedList.AddToEnd(5)

	fmt.Println("Doubly Linked List after adding node 5 to end:")
	doublyLinkedList.IterateThroughList()

	fmt.Println("Adding in AFTER node 5, with new node 6")
	doublyLinkedList.AddAfterThisNode(5, 6)

	fmt.Println("Adding in AFTER node 1, with new node 22")
	doublyLinkedList.AddAfterThisNode(1, 22)

	fmt.Println("Doubly linked list after adding new node 22, after node 1")
	doublyLinkedList.IterateThroughList()

	fmt.Println("Finding node in between node 1 and node 6:", doublyLinkedList.NodeBetweenTwoNodes(1, 6))
}
