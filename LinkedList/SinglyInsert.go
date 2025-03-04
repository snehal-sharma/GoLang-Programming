package main

import "fmt"

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

func (list *LinkedList) insertAtEnd(data int) {
	newNode := &Node{data, nil}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) insertInTheMiddle(data, position int) {
	newNode := &Node{data, nil}
	current := list.head
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

	myList.insertAtStart(45)
	myList.PrintList()

	myList.insertAtStart(11)
	myList.PrintList()

	myList.insertAtEnd(23)
	myList.PrintList()

	myList.insertAtEnd(89)
	myList.PrintList()

	myList.insertInTheMiddle(56, 2)
	myList.PrintList()

}
