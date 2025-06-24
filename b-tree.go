package main

import "fmt"

// How big is the root array?
const rootSize = 4

type Node struct {
	root     [rootSize]*int
	children [rootSize + 1][rootSize]*Node
	isRoot   bool
}

func insert(node *Node, value int) {
	if node == nil {
		return
	}

	var rootLen = len(node.root)
	fmt.Printf("Node size: %d\n", rootLen)
}
