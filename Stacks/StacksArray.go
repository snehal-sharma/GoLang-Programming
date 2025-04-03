package main

import (
	"fmt"
	"math/rand"
)

type Stack struct {
	items    []int
	Capacity int
}

func (s *Stack) IsEmptyStack() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) IsFullStack() bool {
	if len(s.items) == s.Capacity {
		fmt.Println("Stack Overflow")
		return true
	}
	return false
}

func (s *Stack) Push(data int) {
	if !s.IsFullStack() {
		s.items = append(s.items, data)
	}
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) Print() {
	fmt.Println("Printing Stack")
	for _, items := range s.items {
		fmt.Println(items)
	}
}

func main() {
	s := Stack{Capacity: 7}
	fmt.Println("Is Stack empty", s.IsEmptyStack())
	for i := 0; i < 8; i++ {
		s.Push(rand.Intn(100))
	}
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	fmt.Println("Peeking at the top element", s.Peek())
}
