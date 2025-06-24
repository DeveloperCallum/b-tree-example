package main

import (
	"testing"
)

func TestInsertNode(t *testing.T) {
	var node = &Node{
		root:     [4]*int{},
		children: [5][4]*Node{},
		isRoot:   true,
	}

	insert(node, 59)
	insert(node, 23)
	//insert(node, 7)
	//insert(node, 97)
	//
	////Doesn't fit!
	//insert(node, 73)
}
