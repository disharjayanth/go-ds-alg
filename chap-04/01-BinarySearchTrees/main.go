package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

type Tree struct {
	rootNode *Node
}

func (tree *Tree) insert(data int) {
	if tree.rootNode == nil {
		tree.rootNode = &Node{data, nil, nil}
	} else {
		tree.rootNode.insert(data)
	}
}

func (tree *Tree) Search(key int) *Node {
	if key == tree.rootNode.key {
		return tree.rootNode
	}
	return tree.rootNode.Search(key)
}

func (node *Node) insert(data int) {
	if data <= node.key {
		// left
		if node.left == nil {
			node.left = &Node{data, nil, nil}
		} else {
			node.left.insert(data)
		}
	} else {
		// right
		if node.right == nil {
			node.right = &Node{data, nil, nil}
		} else {
			node.right.insert(data)
		}
	}
}

func (node *Node) min() int {
	if node.left == nil {
		return node.key
	}
	return node.left.min()
}

func (node *Node) max() int {
	if node.right == nil {
		return node.key
	}
	return node.right.max()
}

func (node *Node) Search(key int) *Node {
	if key < node.key {
		// left
		if node.left != nil && key == node.left.key {
			return node.left
		} else if node.left != nil {
			return node.left.Search(key)
		} else {
			return nil
		}
	} else {
		// right
		if node.right != nil && key == node.right.key {
			return node.right
		} else if node.right != nil {
			return node.right.Search(key)
		} else {
			return nil
		}
	}
}

func printPreOrder(side string, node *Node) {
	if node == nil {
		return
	}
	fmt.Println(side, node.key)
	printPreOrder("left:", node.left)
	printPreOrder("right:", node.right)
}

func printPostOrder(side string, node *Node) {
	if node == nil {
		return
	}
	printPostOrder("left", node.left)
	printPostOrder("right", node.right)
	fmt.Println(side, node.key)
}

func printInOrder(side string, node *Node) {
	if node == nil {
		return
	}
	printInOrder("left", node.left)
	fmt.Println(side, node.key)
	printInOrder("right", node.right)
}

func main() {
	tree := Tree{}
	tree.insert(8)
	tree.insert(2)
	tree.insert(4)
	tree.insert(1)
	tree.insert(6)
	tree.insert(3)
	tree.insert(7)
	tree.insert(5)

	fmt.Println("PRE - ORDER")
	printPreOrder("root", tree.rootNode)

	fmt.Println("POST - ORDER")
	printPostOrder("root", tree.rootNode)

	fmt.Println("IN - ORDER")
	printInOrder("root", tree.rootNode)

	fmt.Println("Min:", tree.rootNode.min())
	fmt.Println("Max:", tree.rootNode.max())

	fmt.Println("Search for node: 2")
	fmt.Println(*tree.Search(2))

	fmt.Println("Search for node: 6")
	fmt.Println(*tree.Search(6))

	fmt.Println("Search for node: 7")
	fmt.Println(*tree.Search(6))

	fmt.Println("Search for node: 1")
	fmt.Println(*tree.Search(1))

	fmt.Println("Search for node: 10")
	fmt.Println(tree.Search(10))

	fmt.Println("Search for node: -1")
	fmt.Println(tree.Search(-1))
}
