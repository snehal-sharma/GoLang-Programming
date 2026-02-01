// You can edit this code!
// Click here and start typing.
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

func MergeTwoBinaryTrees(root1 *BinaryTreeNode, root2 *BinaryTreeNode) *BinaryTreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Data += root2.Data

	root1.Left = MergeTwoBinaryTrees(root1.Left, root2.Left)
	root1.Right = MergeTwoBinaryTrees(root1.Right, root2.Right)

	return root1
}

func main() {
	binaryTree1 := BinaryTreeNode{Data: 6, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 1}, Right: &BinaryTreeNode{Data: 3}}, Right: &BinaryTreeNode{Data: 8, Left: &BinaryTreeNode{Data: 7}, Right: &BinaryTreeNode{Data: 9}}}
	binaryTree2 := BinaryTreeNode{Data: 1, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 4}, Right: &BinaryTreeNode{Data: 5}}, Right: &BinaryTreeNode{Data: 3, Left: &BinaryTreeNode{Data: 6}, Right: &BinaryTreeNode{Data: 7}}}
	mergedTree := MergeTwoBinaryTrees(&binaryTree1, &binaryTree2)
	PreOrder(mergedTree)

}
