// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func FindMax(root *BinaryTreeNode) int {
	var rootVal, left, right, max int
	if root != nil {
		rootVal = root.Data
		left = FindMax(root.Left)
		right = FindMax(root.Right)

		if left > right {
			max = left
		} else {
			max = right
		}
		if rootVal > max {
			max = rootVal
		}
	}
	return max
}

func FindMin(root *BinaryTreeNode) int {
	var rootVal, left, right, min int
	if root != nil {
		rootVal = root.Data
		left = FindMax(root.Left)
		right = FindMax(root.Right)

		if left < right {
			min = left
		} else {
			min = right
		}
		if rootVal < min {
			min = rootVal
		}
	}
	return min
}

func main() {
	binaryTree := BinaryTreeNode{Data: 1, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 4}, Right: &BinaryTreeNode{Data: 5}}, Right: &BinaryTreeNode{Data: 3, Left: &BinaryTreeNode{Data: 6}, Right: &BinaryTreeNode{Data: 7}}}

	fmt.Printf("Finding Max Element : %d", FindMax(&binaryTree))
	fmt.Printf("\nFinding Min Element : %d", FindMin(&binaryTree))

}
