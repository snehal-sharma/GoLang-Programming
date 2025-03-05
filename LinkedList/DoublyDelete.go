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

func (list *LinkedList) DeleteFromStart() {
	if list.head == nil {
		fmt.Println("Nothing to Delete")
	} else {
		list.head = list.head.next
		list.head.prev = nil
	}

}

func (list *LinkedList) DeleteAtTheEnd() {
	current := list.head
	for current.next != nil {
		current = current.next
	}
	prev := current.prev
	prev.next = nil

}

func (list *LinkedList) DeleteInTheMiddle(pos int) {
	current := list.head
	i := 0
	for current.next != nil && i < pos {
		i++
		current = current.next
	}
	temp := current.prev
	if current.next != nil {
		current.next.prev = temp
	}
	temp.next = current.next

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

	for i := 0; i < 7; i++ {
		myList.insertAtStart(rand.Intn(100))
	}
	myList.PrintList()

	myList.DeleteFromStart()
	myList.PrintList()

	myList.DeleteAtTheEnd()
	myList.PrintList()

	myList.DeleteInTheMiddle(2)
	myList.PrintList()
}
