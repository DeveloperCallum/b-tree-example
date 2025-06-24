package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var node = &Node{
		root:       [4]int{},
		children:   [5][4]*Node{},
		isRoot:     true,
		isModified: false,
	}

	insert(node, 59)
	insert(node, 23)
	insert(node, 7)
	insert(node, 97)

	//Doesn't fit!
	insert(node, 73)
}
