package main

import "fmt"

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func IsBST(root *BinaryTreeNode) int {
	if root == nil {
		return 1
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return 0
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return 0
	}
	if IsBST(root.Left) == 0 || IsBST(root.Right) == 0 {
		return 0
	}
	return 1
}

func main() {
	binaryTree := BinaryTreeNode{Data: 6, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 1}, Right: &BinaryTreeNode{Data: 3}}, Right: &BinaryTreeNode{Data: 8, Left: &BinaryTreeNode{Data: 7}, Right: &BinaryTreeNode{Data: 9}}}
	val := IsBST(&binaryTree)
	fmt.Println("Is Binary Tree", val)

}
