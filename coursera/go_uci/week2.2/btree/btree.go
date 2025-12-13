package main

import (
	"fmt"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Left: nil, Right: nil}
	}
	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	}
	return root
}

func search(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if root.Key == key {
		return true
	}
	if key < root.Key {
		return search(root.Left, key)
	}
	return search(root.Right, key)
}

func main() {
	var root *Node
	keys := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, key := range keys {
		root = insert(root, key)
	}
	fmt.Println("Binary Tree Example:")

	searchKey := 6
	if search(root, searchKey) {
		fmt.Printf("Key %d found in the binary tree.\n", searchKey)
	} else {
		fmt.Printf("Key %d not found in the binary tree.\n", searchKey)
	}
}
