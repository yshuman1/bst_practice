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
