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

func (list *LinkedList) FindNthNodeFromEndOfList(pos, len int) int {
	i := 0
	if pos > len {
		fmt.Println("\nFewer Number of nodes in the list")
		return -1
	}
	nodePos := len - pos
	current := list.head
	for current.next != nil && i < nodePos {
		current = current.next
		i++
	}
	return current.data
}

func (list *LinkedList) PrintList() int {
	count := 0
	fmt.Println("\nPrinting List")
	current := list.head
	for current != nil {
		fmt.Printf("\t %d", current.data)
		current = current.next
		count++
	}
	return count
}

func main() {
	myList := LinkedList{}

	for i := 0; i < 10; i++ {
		myList.insertAtStart(rand.Intn(100))
	}
	length := myList.PrintList()
	fmt.Printf("\nLength of List is %v", length)

	pos := rand.Intn(10)
	node := myList.FindNthNodeFromEndOfList(pos, length)
	if node != -1 {
		fmt.Printf("\n%vth Node from the End of the List is %v", pos, node)
	}

}
