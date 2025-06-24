package main

import "fmt"

// How big is the root array?
const rootSize = 4

type Node struct {
	root       [rootSize]int
	children   [rootSize + 1][rootSize]*Node
	isRoot     bool
	isModified bool
}

func insert(node *Node, value int) {
	var rootLen = len(node.root)
	fmt.Printf("Node size: %d\n", rootLen)
}

func addValueToRoot(node *Node, value int) {

	//Exception to b-tree rule, the first code can have 1 element.
	if node.isRoot && !node.isModified {
		node.root[0] = value
		node.isModified = true
		return
	}
	
}
