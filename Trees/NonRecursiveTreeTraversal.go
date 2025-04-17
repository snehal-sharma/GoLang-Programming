package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type Stack struct {
	items    []*BinaryTreeNode
	Capacity int
}

func (s *Stack) CreateStack() {
	s.items = make([]*BinaryTreeNode, 0, s.Capacity)
}

func (s *Stack) IsEmptyStack() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) IsFullStack() bool {
	if len(s.items) == s.Capacity {
		fmt.Println("Stack Overflow")
		return true
	}
	return false
}

func (s *Stack) Push(root *BinaryTreeNode) {
	if !s.IsFullStack() {
		s.items = append(s.items, root)
	}
}

func (s *Stack) Pop() *BinaryTreeNode {
	root := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return root
}

func (s *Stack) Peek() *BinaryTreeNode {
	return s.items[len(s.items)-1]
}

func (s *Stack) Print() {
	fmt.Println("Printing Stack")
	for _, items := range s.items {
		fmt.Println(items)
	}
}

func NonRecursivePreOrderTraversal(root *BinaryTreeNode) []int {
	myStack := Stack{Capacity: 8}
	myStack.CreateStack()
	var traversal []int
	for {
		for root != nil {
			traversal = append(traversal, root.Data)
			myStack.Push(root)
			myStack.Print()
			root = root.Left
		}
		if myStack.IsEmptyStack() {
			break
		}
		root = myStack.Pop()
		myStack.Print()
		root = root.Right
	}
	return traversal
}

func NonRecursiveInOrderTraversal(root *BinaryTreeNode) []int {
	myStack := Stack{Capacity: 8}
	myStack.CreateStack()
	var traversal []int
	for {
		for root != nil {
			myStack.Push(root)
			myStack.Print()
			root = root.Left
		}
		if myStack.IsEmptyStack() {
			break
		}
		root = myStack.Pop()
		myStack.Print()
		traversal = append(traversal, root.Data)
		root = root.Right
	}
	return traversal
}

func NonRecursivePostOrderTraversal(root *BinaryTreeNode) []int {
	myStack := Stack{Capacity: 8}
	myStack.CreateStack()
	var traversal []int
	var previous *BinaryTreeNode
	for {
		for root != nil {
			myStack.Push(root)
			myStack.Print()
			root = root.Left
		}
		for root == nil && !myStack.IsEmptyStack() {
			root = myStack.Peek()
			if root.Right == nil || root.Right == previous {
				traversal = append(traversal, root.Data)
				_ = myStack.Pop()
				myStack.Print()
				previous = root
				root = nil
			} else {
				root = root.Right
			}
		}
		if myStack.IsEmptyStack() {
			break
		}
	}
	return traversal
}

func main() {
	binaryTree := BinaryTreeNode{Data: 1, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 4}, Right: &BinaryTreeNode{Data: 5}}, Right: &BinaryTreeNode{Data: 3, Left: &BinaryTreeNode{Data: 6}, Right: &BinaryTreeNode{Data: 7}}}
	fmt.Println(binaryTree)

	traversal := NonRecursivePreOrderTraversal(&binaryTree)
	fmt.Println("Non Recursive PreOrder Traversal :", traversal)
	fmt.Printf("\n\n")

	traversal = NonRecursiveInOrderTraversal(&binaryTree)
	fmt.Println("Non Recursive InOrder Traversal :", traversal)
	fmt.Printf("\n\n")

	traversal = NonRecursivePostOrderTraversal(&binaryTree)
	fmt.Println("Non Recursive PostOrder Traversal :", traversal)

}
