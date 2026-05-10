package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
	Height int
}

func SaveHeight(value int) int {
	node = append(node, value)
	return node
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}
		
	return node
}

func main() {

	var root *Node

	arr := []int{12, 45, 23, 67, 98, 34, 7, 3}

	for index, value := range arr {
		root = insert(root, value)
		SaveHeight(index)

		fmt.Println("\n\n root ", root)
	}

	// fmt.Println("\n\n root ", root)
}