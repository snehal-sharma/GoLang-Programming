// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
)

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func FindElementBinaryTreeRecursion(root *BinaryTreeNode, data int) int {
	if root == nil {
		return 0
	} else {
		if root.Data == data {
			return 1
		} else {
			temp := FindElementBinaryTreeRecursion(root.Left, data)
			if temp != 0 {
				return temp
			} else {
				return FindElementBinaryTreeRecursion(root.Right, data)
			}
		}
	}
}

func main() {
	binaryTree := BinaryTreeNode{Data: 1, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 4}, Right: &BinaryTreeNode{Data: 5}}, Right: &BinaryTreeNode{Data: 3, Left: &BinaryTreeNode{Data: 6}, Right: &BinaryTreeNode{Data: 7}}}
	element := rand.Intn(10)
	fmt.Printf("Finding Element : %d", element)
	result := FindElementBinaryTreeRecursion(&binaryTree, element)
	if result == 1 {
		fmt.Printf("\nFound Element : %d", element)
	} else {
		fmt.Printf("\nElement %d Not Found", element)
	}

}
