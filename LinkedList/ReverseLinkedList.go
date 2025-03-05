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

func (list *LinkedList) ReverseList() {
	var prev, current, head *Node
	head = list.head
	for head != nil {
		current = head.next
		head.next = prev
		prev = head
		head = current
	}
	list.head = prev
}

func (list *LinkedList) PrintList() {
	fmt.Println("\nPrinting List")
	current := list.head
	for current != nil {
		fmt.Printf("\t %d", current.data)
		current = current.next
	}
}

func main() {
	myList := LinkedList{}

	for i := 0; i < 7; i++ {
		myList.insertAtStart(rand.Intn(100))
	}
	myList.PrintList()

	myList.ReverseList()
	myList.PrintList()

}
