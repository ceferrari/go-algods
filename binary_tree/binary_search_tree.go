package binary_tree

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (n *Node) insert(val int) {
	if val < n.val {
		if n.left == nil {
			n.left = &Node{val: val}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val: val}
		} else {
			n.right.insert(val)
		}
	}
}

func (b *BinarySearchTree) Insert(val int) {
	if b.root == nil {
		b.root = &Node{val: val}
	} else {
		b.root.insert(val)
	}
	fmt.Print(val, " ")
}

func preOrder(node *Node, function func(int)) {
	if node == nil {
		return
	}
	function(node.val)
	preOrder(node.left, function)
	preOrder(node.right, function)
}

func (b *BinarySearchTree) PreOrder(function func(int)) {
	preOrder(b.root, function)
}

func inOrder(node *Node, function func(int)) {
	if node == nil {
		return
	}
	inOrder(node.left, function)
	function(node.val)
	inOrder(node.right, function)
}

func (b *BinarySearchTree) InOrder(function func(int)) {
	inOrder(b.root, function)
}

func postOrder(node *Node, function func(int)) {
	if node == nil {
		return
	}
	postOrder(node.left, function)
	postOrder(node.right, function)
	function(node.val)
}

func (b *BinarySearchTree) PostOrder(function func(int)) {
	postOrder(b.root, function)
}

func Traverse() {
	b := &BinarySearchTree{}
	b.Insert(25)
	b.Insert(20)
	b.Insert(36)
	b.Insert(10)
	b.Insert(22)
	b.Insert(30)
	b.Insert(40)
	b.Insert(5)
	b.Insert(12)
	b.Insert(28)
	b.Insert(38)
	b.Insert(48)
	b.Insert(1)
	b.Insert(8)
	b.Insert(15)
	b.Insert(45)
	b.Insert(50)
	fmt.Println()

	function := func(val int) {
		fmt.Print(val, " ")
	}

	b.PreOrder(function)
	fmt.Println("(pre)")
	b.InOrder(function)
	fmt.Println("(in)")
	b.PostOrder(function)
	fmt.Println("(post)")
}
