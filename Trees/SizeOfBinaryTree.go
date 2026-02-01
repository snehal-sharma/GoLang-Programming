package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type Queue struct {
	Rear, Front, Data, Capacity int
	arr                         []*BinaryTreeNode
}

func (myQueue *Queue) CreateQueue() {
	myQueue.arr = make([]*BinaryTreeNode, 0, myQueue.Capacity)
}

func (myQueue *Queue) IsEmptyQueue() bool {
	if myQueue.Front == -1 {
		return true
	}
	return false
}
func (myQueue *Queue) IsFullQueue() bool {
	if myQueue.Front == myQueue.Capacity {
		return true
	}
	return false
}

func (myQueue *Queue) EnQueue(root *BinaryTreeNode) {
	if !myQueue.IsFullQueue() {
		myQueue.arr = append(myQueue.arr, root)
		if myQueue.Rear == -1 {
			myQueue.Rear = 0
			myQueue.Front = 0
		}
		myQueue.Front++
	} else {
		fmt.Println("Queue Overflow: Trying to Insert in a full Queue")
	}
}

func (myQueue *Queue) DeQueue() *BinaryTreeNode {
	if !myQueue.IsEmptyQueue() {
		myQueue.Rear++
		if myQueue.Rear == myQueue.Front {
			myQueue.Rear = -1
			myQueue.Front = -1
		}
		root := myQueue.arr[0]
		myQueue.arr = myQueue.arr[1:]
		return root
	} else {
		fmt.Println("Queue Underflow:Deleting from an empty Queue")
		return nil
	}

}

func SizeOfBinaryTreeRecursive(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	} else {
		return (SizeOfBinaryTreeRecursive(root.Left) + 1 + SizeOfBinaryTreeRecursive(root.Right))
	}
}

func SizeOfBinaryTree(root *BinaryTreeNode) int {
	myQueue := Queue{Front: -1, Rear: -1, Capacity: 8}
	myQueue.CreateQueue()
	if root == nil {
		return 0
	}
	count := 0
	myQueue.EnQueue(root)
	for !myQueue.IsEmptyQueue() {
		root = myQueue.DeQueue()
		count++
		if root.Left != nil {
			myQueue.EnQueue(root.Left)
		}
		if root.Right != nil {
			myQueue.EnQueue(root.Right)
		}
	}
	return count

}

func main() {
	binaryTree := BinaryTreeNode{Data: 1, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 4}, Right: &BinaryTreeNode{Data: 5}}, Right: &BinaryTreeNode{Data: 3, Left: &BinaryTreeNode{Data: 6}, Right: &BinaryTreeNode{Data: 7}}}
	size := SizeOfBinaryTreeRecursive(&binaryTree)
	fmt.Println("Size of Binary Tree :", size)
}
