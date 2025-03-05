package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtStart(data int) {
	newNode := &Node{data, nil, nil}
	newNode.next = list.head
	if list.head != nil {
		list.head.prev = newNode
	}
	list.head = newNode
}

func (list *LinkedList) insertAtEnd(data int) {
	newNode := &Node{data, nil, nil}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	newNode.prev = current
	current.next = newNode
}

func (list *LinkedList) insertInTheMiddle(data, position int) {
	newNode := &Node{data, nil, nil}
	current := list.head
	i := 0
	for current.next != nil && i < position-1 {
		i++
		current = current.next
	}

	newNode.next = current.next
	newNode.prev = current
	if current.next != nil {
		current.next.prev = newNode
	}
	current.next = newNode
}

func (list *LinkedList) PrintList() {
	fmt.Println("\nPrinting List")
	current := list.head
	for current != nil {
		fmt.Printf("\t %v", current)
		current = current.next
	}
}

func main() {
	myList := LinkedList{}

	myList.insertAtStart(rand.Intn(100))
	myList.PrintList()

	myList.insertAtStart(rand.Intn(100))
	myList.PrintList()

	myList.insertAtEnd(rand.Intn(100))
	myList.PrintList()

	myList.insertAtEnd(rand.Intn(100))
	myList.PrintList()

	myList.insertInTheMiddle(rand.Intn(100), 2)
	myList.PrintList()
}
