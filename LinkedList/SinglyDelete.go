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

func (list *LinkedList) DeleteFromStart() {
	if list.head == nil {
		fmt.Println("Nothing to Delete")
	} else {
		list.head = list.head.next
	}

}

func (list *LinkedList) DeleteAtTheEnd() {
	current := list.head
	prev := current
	for current.next != nil {
		prev = current
		current = current.next
	}
	prev.next = nil

}

func (list *LinkedList) DeleteInTheMiddle(pos int) {
	current := list.head
	prev := current
	i := 0
	for current.next != nil && i < pos {
		i++
		prev = current
		current = current.next
	}
	prev.next = current.next

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

	myList.DeleteFromStart()
	myList.PrintList()

	myList.DeleteAtTheEnd()
	myList.PrintList()

	myList.DeleteInTheMiddle(3)
	myList.PrintList()
}
