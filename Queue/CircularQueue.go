
package main

import (
	"fmt"
	"math/rand"
)

type Queue struct {
	Rear, Front, Capacity int
	arr                   []int
}

func (myQueue *Queue) IsEmptyQueue() bool {
	if myQueue.Front == -1 {
		return true
	}
	return false
}

func (myQueue *Queue) IsFullQueue() bool {
	if (myQueue.Rear+1)%myQueue.Capacity == myQueue.Front {
		return true
	}
	return false
}

func (myQueue *Queue) EnQueue(data int) {
	if !myQueue.IsFullQueue() {
		myQueue.Rear = (myQueue.Rear + 1) % myQueue.Capacity
		myQueue.arr[myQueue.Rear] = data
		if myQueue.Front == -1 {
			myQueue.Front = myQueue.Rear
		}
	} else {
		fmt.Println("Queue Overflow: Trying to Insert in a full Queue")
	}
}

func (myQueue *Queue) DeQueue() {
	if myQueue.IsEmptyQueue() {
		fmt.Println("Queue Underflow:Deleting from an empty Queue")
		return
	} else {
		myQueue.arr[myQueue.Front] = 0
		if myQueue.Front == myQueue.Rear {
			myQueue.Front = -1
			myQueue.Rear = -1
		} else {
			myQueue.Front = (myQueue.Front + 1) % myQueue.Capacity
		}
	}

}

func main() {
	myQueue := Queue{Rear: -1, Front: -1, Capacity: 7}
	myQueue.arr = make([]int, myQueue.Capacity)
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

	for i := 0; i < 7; i++ {
		myQueue.EnQueue(rand.Intn(100))
		fmt.Println("After EnQueue")
		fmt.Println(myQueue)
	}

	for i := 0; i < 9; i++ {
		myQueue.DeQueue()
		fmt.Println("After DeQueue")
		fmt.Println(myQueue)
	}

}
