package main

import "fmt"

type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

func (tree *Tree) insert(m int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{
				LeftNode:  nil,
				Value:     m,
				RightNode: nil,
			}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{
					LeftNode:  nil,
					Value:     m,
					RightNode: nil,
				}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{
			LeftNode:  nil,
			Value:     m,
			RightNode: nil,
		}
	}
}

func print(tree *Tree) {
	if tree != nil {
		fmt.Println("*Value:", tree.Value)
		fmt.Println("<- Left Node:")
		print(tree.LeftNode)
		fmt.Println("-> Right Node:")
		print(tree.RightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

func main() {
	var tree *Tree = &Tree{
		LeftNode:  nil,
		Value:     1,
		RightNode: nil,
	}
	fmt.Println("----------------------------")
	fmt.Println("Insert 1")
	print(tree)

	fmt.Println("----------------------------")
	fmt.Println("Insert 3")
	tree.insert(3)
	print(tree)

	fmt.Println("----------------------------")
	fmt.Println("Insert 5")
	tree.insert(5)
	print(tree)

	fmt.Println("----------------------------")
	fmt.Println("Insert 7")
	tree.insert(7)
	print(tree)
}
