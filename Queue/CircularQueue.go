
package main

import (
	"fmt"
	"math/rand"
)

type Queue struct {
	Rear, Front, Data, Capacity int
	arr                         []int
}

func (myQueue *Queue) IsEmptyQueue() bool {
	if myQueue.Front == -1 {
		return true
	}
	return false
}

func (myQueue *Queue) IsFullQueue() bool {
	if myQueue.Front == myQueue.Capacity {
		return true
	}
	return false
}

func (myQueue *Queue) EnQueue(data int) {
	if !myQueue.IsFullQueue() {
		myQueue.arr = append(myQueue.arr, data)
		if myQueue.Rear == -1 {
			myQueue.Rear = 0
			myQueue.Front = 0
		}
		myQueue.Front++
	} else {
		fmt.Println("Queue Overflow: Trying to Insert in a full Queue")
	}
}

func (myQueue *Queue) DeQueue() {
	if myQueue.Rear == myQueue.Front {
		myQueue.Rear = -1
		myQueue.Front = -1
	}
	if !myQueue.IsEmptyQueue() {
		myQueue.Rear++
		myQueue.arr = myQueue.arr[1:]
	} else {
		fmt.Println("Queue Underflow:Deleting from an empty Queue")
		return
	}

}

func main() {
	myQueue := Queue{Rear: -1, Front: -1, Capacity: 7}
	myQueue.arr = make([]int, 0, myQueue.Capacity)
	fmt.Println(myQueue)
	fmt.Println(myQueue.IsEmptyQueue())

	for i := 0; i < 4; i++ {
		myQueue.EnQueue(rand.Intn(100))
		fmt.Println("After EnQueue")
		fmt.Println(myQueue)
	}

	for i := 0; i < 2; i++ {
		myQueue.DeQueue()
		fmt.Println("After DeQueue")
		fmt.Println(myQueue)
	}

	for i := 0; i < 4; i++ {
		myQueue.EnQueue(rand.Intn(100))
		fmt.Println("After EnQueue")
		fmt.Println(myQueue)
	}

	for i := 0; i < 6; i++ {
		myQueue.DeQueue()
		fmt.Println("After DeQueue")
		fmt.Println(myQueue)
	}

}
