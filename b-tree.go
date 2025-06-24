package main

import "fmt"

// How big is the root array?
const rootSize = 4

type Node struct {
	root     [rootSize]*int
	children [rootSize + 1][rootSize]*Node
	isRoot   bool
}

func (n *Node) insert(value int) {
	//Each node must have AT LEAST 2 nodes.
	if n.isRoot {
		if n.root[0] == nil {
			n.root[0] = &value //Start node can have one element.
			return
		}

		if n.root[1] == nil {
			n.root[0] = &value //Start node can have one element.
			return
		}
	}

	var rootLen = len(n.root)
	fmt.Printf("Node size: %d\n", rootLen)
}

func (n *Node) hasSpaceInRoot() bool {
	for i := 0; i < len(n.root); i++ {
		if n.root[i] == nil {
			return true
		}
	}

	return false
}

func (n *Node) insertIntoRootArr(value int) {
	if n.hasSpaceInRoot() {
		for i := 0; i < len(n.root); i++ {
			arrValue := n.root[i] //5

			if value >= *arrValue { //insert

			}
		}
	}
}

func (n *Node) shift(value int) {
	for i := len(n.root) - 2; i >= 0; i-- {
		node := n.root[i]
		lastCheckedNode := n.root[i+1]

		if node != nil {
			insert := *node <= value

			if lastCheckedNode == nil {
				n.root[i+1] = node
				n.root[i] = nil
			}

			if insert { //Node should be smaller if we insert it like 1, 2, node, value, 5
				n.root[i] = &value
				break
			}
		}
	}
}
