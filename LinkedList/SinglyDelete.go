package main

import (
	"fmt"
	"math/rand"
)

type LinkedList struct {
	data int
	next *LinkedList
}

func insertAtStart(data int, head *LinkedList) *LinkedList {
	newNode := LinkedList{data, nil}
	newNode.next = head
	head = &newNode
	return head
}

func DeleteFromStart(head *LinkedList) *LinkedList {
	if head == nil {
		fmt.Println("Nothing to Delete")
	} else {
		head = head.next
	}
	return head

}

func DeleteAtTheEnd(head *LinkedList) {
	current := head
	prev := current
	for current.next != nil {
		prev = current
		current = current.next
	}
	prev.next = nil

}

func DeleteInTheMiddle(pos int, head *LinkedList) {
	current := head
	prev := current
	i := 0
	for current.next != nil && i < pos {
		i++
		prev = current
		current = current.next
	}
	prev.next = current.next

}

func PrintList(head *LinkedList) {
	fmt.Println("\nPrinting List")
	current := head
	for current != nil {
		fmt.Printf("\t %d", current.data)
		current = current.next
	}
}

func main() {
	//	var head *LinkedList
	//	DeleteFromStart(head)
	head := &LinkedList{10, nil}
	PrintList(head)
	for i := 0; i < 7; i++ {
		head = insertAtStart(rand.Intn(100), head)
	}
	PrintList(head)
	head = DeleteFromStart(head)
	PrintList(head)
	DeleteAtTheEnd(head)
	PrintList(head)
	DeleteInTheMiddle(3, head)
	PrintList(head)

}
