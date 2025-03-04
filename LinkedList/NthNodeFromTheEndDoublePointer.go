package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtStart(data int) {
	newNode := &Node{data, nil}
	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) FindNthNodeFromEndOfList(pos int) int {
	i := 0
	currentNthNode := list.head
	current := list.head
	for currentNthNode.next != nil && i < pos-1 {
		fmt.Printf("\t %v", currentNthNode.data)
		currentNthNode = currentNthNode.next
		i++
	}
	if currentNthNode.next == nil {
		fmt.Println("\nFewer items in the linked list")
		return -1
	}
	for currentNthNode.next != nil {
		current = current.next
		currentNthNode = currentNthNode.next
	}

	return current.data
}

func (list *LinkedList) PrintList() {
	fmt.Println("\nPrinting List")
	current := list.head
	for current != nil {
		fmt.Printf("\t %v", current.data)
		current = current.next
	}
}

func main() {
	myList := LinkedList{}

	for i := 0; i < 10; i++ {
		myList.insertAtStart(rand.Intn(100))
	}
	myList.PrintList()

	pos := rand.Intn(10)
	fmt.Printf("\nFind %vth Node from the End of the List\n", pos)
	node := myList.FindNthNodeFromEndOfList(pos)
	if node != -1 {
		fmt.Printf("\n%vth Node from the End of the List is %v", pos, node)
	}

}
