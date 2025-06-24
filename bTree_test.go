package main

import (
	"testing"
)

func TestCheckSift(t *testing.T) {
	arr := [4]*int{ //Create a new array, take the pass the memory addresses of the value from an anonymous function.
		func() *int { v := 4; return &v }(),
		func() *int { v := 5; return &v }(),
	}

	var node = &Node{
		root:     arr,
		children: [5][4]*Node{},
		isRoot:   true,
	}

	node.shift(5)
	node.shift(7)
}

func TestInsertNode(t *testing.T) {
	var node = &Node{
		root:     [4]*int{},
		children: [5][4]*Node{},
		isRoot:   true,
	}

	node.insert(59)
	node.insert(23)
	node.insert(7)
	//insert(node, 97)
	//
	////Doesn't fit!
	//insert(node, 73)
}
