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

func (list *LinkedList) FloydAlgorithm() bool {
	var slowPtr, fastPtr *Node
	slowPtr = list.head
	fastPtr = list.head
	for slowPtr != nil && fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			return true
		}

	}
	return false
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

	for i := 0; i < 10; i++ {
		myList.insertAtStart(rand.Intn(100))
	}
	myList.PrintList()

	fmt.Println("\nThe LinkedList Contains a Loop :", myList.FloydAlgorithm())

}
