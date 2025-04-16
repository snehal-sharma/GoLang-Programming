
package main

import "fmt"

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func PreOrder(root *BinaryTreeNode) {
	if root != nil {
		fmt.Println(root.Data)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

func main() {
	binaryTree := BinaryTreeNode{Data: 1, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 4}, Right: &BinaryTreeNode{Data: 5}}, Right: &BinaryTreeNode{Data: 3, Left: &BinaryTreeNode{Data: 6}, Right: &BinaryTreeNode{Data: 7}}}
	PreOrder(&binaryTree)
}
