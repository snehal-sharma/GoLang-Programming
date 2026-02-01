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

func LevelOrderTraversalNonRecursive(root *BinaryTreeNode) []int {
	myQueue := Queue{Rear: -1, Front: -1, Capacity: 8}
	myQueue.CreateQueue()
	myQueue.EnQueue(root)
	var traversal []int
	for !myQueue.IsEmptyQueue() {
		temp := myQueue.DeQueue()
		traversal = append(traversal, temp.Data)
		if temp.Left != nil {
			myQueue.EnQueue(temp.Left)
		}
		if temp.Right != nil {
			myQueue.EnQueue(temp.Right)
		}
	}
	return traversal
}

func InsertIntoBinaryTree(root *BinaryTreeNode) {
	myQueue := Queue{Front: -1, Rear: -1, Capacity: 8}
	myQueue.CreateQueue()
	newNode := &BinaryTreeNode{Data: 22, Left: nil, Right: nil}
	if root == nil {
		root = newNode
		return
	}
	myQueue.EnQueue(root)
	for !myQueue.IsEmptyQueue() {
		root = myQueue.DeQueue()
		if root.Left != nil {
			myQueue.EnQueue(root.Left)
		} else {
			root.Left = newNode
			return
		}
		if root.Right != nil {
			myQueue.EnQueue(root.Right)
		} else {
			root.Right = newNode
			return
		}
	}

}

func main() {
	binaryTree := BinaryTreeNode{Data: 1, Left: &BinaryTreeNode{Data: 2, Left: &BinaryTreeNode{Data: 4}, Right: &BinaryTreeNode{Data: 5}}, Right: &BinaryTreeNode{Data: 3, Left: &BinaryTreeNode{Data: 6}, Right: &BinaryTreeNode{Data: 7}}}
	InsertIntoBinaryTree(&binaryTree)
	trav := LevelOrderTraversalNonRecursive(&binaryTree)
	fmt.Println(trav)
}
