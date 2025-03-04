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

func (list *LinkedList) PrintList() (int, map[int]*Node) {
	count := 0
	hashMap := make(map[int]*Node)
	fmt.Println("\nPrinting List")
	current := list.head
	for current != nil {
		hashMap[count] = current
		fmt.Printf("\t %d", current.data)
		current = current.next
		count++
	}
	return count, hashMap
}

func main() {
	myList := LinkedList{}

	for i := 0; i < 10; i++ {
		myList.insertAtStart(rand.Intn(100))
	}
	length, hashMap := myList.PrintList()
	fmt.Printf("\nLength of List is %v", length)

	pos := rand.Intn(10)
	nodePos := length - pos
	node, ok := hashMap[nodePos]
	if ok {
		fmt.Printf("\n%vth Node from the End of the List is %v", pos, node.data)
	}

}
