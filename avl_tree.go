package main 

import (
	"fmt"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
	Height uint
}

func GetHeight(n *Node) uint {
	if n == nil {
		return 0
	}

	return n.Height
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	} else {
		return node
	}
	return node
}

func main() {
	var root *Node
	arr := []int{12, 13, 17, 11, 28, 2, 39, 3}

	for _, value := range arr {
		// fmt.Printf("\n index %d \n value %d ", index, value)
		
		root = insert(root, value)

		fmt.Printf("Inserted %d -> root is now %d\n", value, root)
	} 

}


