package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = node
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node)
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) AddToLast(property int) {
	newNode := Node{}
	newNode.property = property
	newNode.nextNode = nil

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			node.nextNode = &newNode
			return
		}
	}

	// Or use
	// lastNode := linkedList.LastNode()
	// if lastNode != nil {
	// 	lastNode.nextNode = &newNode
	// }
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			return node
		}
	}
	return nil
}

func (linked *LinkedList) AddAfterThisNode(thisNode int, property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	nodeWith := linked.NodeWithValue(thisNode)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	linkedList := LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(4)
	linkedList.AddToHead(8)
	linkedList.AddToHead(12)

	linkedList.IterateList()

	lastNode := linkedList.LastNode()
	if lastNode != nil {
		fmt.Println("Last node:", *lastNode)
	}

	fmt.Println("Add 16 to last node")
	linkedList.AddToLast(16)

	fmt.Println("Iterating again after adding 16 to last node:")
	linkedList.IterateList()

	matchNode := linkedList.NodeWithValue(12)
	fmt.Println("Check if node has value of 12:", matchNode)

	fmt.Println("Add after 8 node, node to be added 10")
	linkedList.AddAfterThisNode(8, 10)

	fmt.Println("Iterating after adding:")
	linkedList.IterateList()
}
