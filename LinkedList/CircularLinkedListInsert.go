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
	newNode.next = newNode
	current := list.head
	if list.head == nil {
		list.head = newNode
	} else {
		for current.next != list.head {
			current = current.next
		}
		newNode.next = list.head
		current.next = newNode
		list.head = newNode
	}
}

func (list *LinkedList) insertAtEnd(data int) {
	newNode := &Node{data, nil}
	newNode.next = newNode
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != list.head {
			current = current.next
		}
		newNode.next = list.head
		current.next = newNode
	}
}

func (list *LinkedList) PrintList() {
	fmt.Println("\nPrinting List")
	if list.head == nil {
		fmt.Println("Circular Linked List is Empty")
	}
	current := list.head
	fmt.Printf("\t %v", current)
	current = current.next
	for current != list.head {
		fmt.Printf("\t %v", current)
		current = current.next
	}
}

func main() {
	myList := LinkedList{}
	for i := 0; i < 3; i++ {
		myList.insertAtStart(rand.Intn(100))
		myList.PrintList()
	}

	for i := 0; i < 3; i++ {
		myList.insertAtEnd(rand.Intn(100))
		myList.PrintList()
	}

}
