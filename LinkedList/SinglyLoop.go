
package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Data int
	Next *Node
}

func PrintList(head *Node) {
	ptrNode := head
	for ; ptrNode.Next != nil; ptrNode = ptrNode.Next {
		fmt.Println(ptrNode)
	}
}

func main() {
	var head, prev *Node
	head = &Node{Data: 12, Next: nil}
	prev = head
	for i := 0; i < 10; i++ {
		newNode := Node{Data: rand.Intn(10), Next: nil}
		prev.Next = &newNode
		prev = &newNode
	}
	PrintList(head)
	fmt.Println("Hello, 世界")
}
