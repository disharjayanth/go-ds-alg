package main

import (
	"fmt"
)

type BinaryNode struct {
	left  *BinaryNode
	value int
	right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}

func (bt *BinaryTree) insert(val int) {
	if bt.root == nil {
		bt.root = &BinaryNode{nil, val, nil}
	} else {
		bt.root.insert(val)
	}
}

func (bn *BinaryNode) insert(val int) {
	if bn == nil {
		return
	} else if val <= bn.value {
		// left side
		if bn.left == nil {
			bn.left = &BinaryNode{nil, val, nil}
		} else {
			bn.left.insert(val)
		}
	} else {
		// right side
		if bn.right == nil {
			bn.right = &BinaryNode{nil, val, nil}
		} else {
			bn.right.insert(val)
		}
	}
}

// func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
// 	if node == nil {
// 		return
// 	}

// 	for i := 0; i < ns; i++ {
// 		fmt.Fprint(w, " ")
// 	}
// 	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
// 	print(w, node.left, ns+2, 'L')
// 	print(w, node.right, ns+2, 'R')
// }

func print(bn *BinaryNode) {
	if bn != nil {
		fmt.Println("Value:", bn.value)
		fmt.Println("Tree Node Left:")
		print(bn.left)
		fmt.Println("Tree Right Node:")
		print(bn.right)
	} else {
		fmt.Println("Nil")
	}
}

func main() {
	tree := &BinaryTree{}
	tree.insert(100)
	tree.insert(-20)
	tree.insert(-50)
	tree.insert(-15)
	tree.insert(-60)
	tree.insert(50)
	tree.insert(60)
	tree.insert(55)
	tree.insert(85)
	tree.insert(15)
	tree.insert(5)
	tree.insert(-10)

	print(tree.root)
	// print(os.Stdout, tree.root, 0, 'M')
}
