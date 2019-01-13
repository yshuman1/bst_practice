package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func main() {
	tree := &Node{6, nil, nil}
	fmt.Println(tree.Key)
}

func (n *Node) Search(key int) bool {
	// if the node is nil, the key does not exist
	if n == nil {
		return false
	}
	//if the key is greater than the node we are at move to the right
	if n.Key < key {
		//move right
		return n.Right.Search(key)
	} else if n.Key > key {
		//if the key is less than the node we are at move to the left
		//move left
		return n.Left.Search(key)
	}
	//you found the key if it is equal to your current node
	return true
}

func (n *Node) Insert(key int) {
	// if the key is greater than the node you are at and the right node is nil, insert. Otherwise,keep searching to the right.
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}

		} else {
			n.Left.Insert(key)
		}
	}
}
