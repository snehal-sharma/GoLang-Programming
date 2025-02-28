package main

import "fmt"

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

func insertAtEnd(data int, head *LinkedList) {
	newNode := &LinkedList{data, nil}
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func insertInTheMiddle(data, position int, head *LinkedList) {
	newNode := &LinkedList{data, nil}
	current := head
	prev := current
	i := 0
	for current.next != nil && i < position {
		i++
		prev = current
		current = current.next
	}
	prev.next = newNode
	newNode.next = current
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
	head := &LinkedList{3, nil}
	PrintList(head)
	head = insertAtStart(45, head)
	PrintList(head)
	insertAtEnd(23, head)
	PrintList(head)
	insertAtEnd(89, head)
	PrintList(head)
	insertInTheMiddle(56, 2, head)
	PrintList(head)
}
