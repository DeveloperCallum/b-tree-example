package main

import (
	"errors"
	"fmt"
)

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
			n.root[1] = &value
			return
		}
	}

	if n.hasSpaceInRoot() { //Insert value into the array if the array has space.
		_ = n.shift(value)
	}

	//var rootLen = len(n.root)
	fmt.Println(n.root)
}

func (n *Node) hasSpaceInRoot() bool {
	for i := 0; i < len(n.root); i++ {
		if n.root[i] == nil {
			return true
		}
	}

	return false
}

/*
Shift the array and put the value in the array.
*/
func (n *Node) shift(value int) error {
	if !n.hasSpaceInRoot() {
		return nil
	}

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

	//MUST FULL
	return errors.New("failed to insert value")
}

func (n *Node) fits(value int) bool {
	for i := len(n.root) - 2; i >= 0; i-- {
		node := n.root[i]
		insert := *node <= value

		if insert {
			return true
		}
	}

	return false
}
