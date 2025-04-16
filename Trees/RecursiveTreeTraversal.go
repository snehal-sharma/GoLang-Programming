package main

import "fmt"

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func PreOrder(root *BinaryTreeNode) {
	if root != nil {
		fmt.Printf("\t%d", root.Data)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

func InOrder(root *BinaryTreeNode) {
	if root != nil {
		InOrder(root.Left)
		fmt.Printf("\t%d", root.Data)
		InOrder(root.Right)
	}
}

func PostOrder(root *BinaryTreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Printf("\t%d", root.Data)
	}
}

func main() {
	binaryTree := BinaryTreeNode{Data: 1, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 4}, Right: &BinaryTreeNode{Data: 5}}, Right: &BinaryTreeNode{Data: 3, Left: &BinaryTreeNode{Data: 6}, Right: &BinaryTreeNode{Data: 7}}}
	fmt.Println("PreOrder Tree Traversal : ")
	PreOrder(&binaryTree)

	fmt.Println("\nInOrder Tree Traversal : ")
	InOrder(&binaryTree)

	fmt.Println("\nPostOrder Tree Traversal : ")
	PostOrder(&binaryTree)
}
