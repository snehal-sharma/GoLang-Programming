package main

import (
	"fmt"
	"sync"
)

func MergeChannel(left chan int, right chan int, arrayc chan int) {
	val1, ok1 := <-left
	val2, ok2 := <-right
	defer close(arrayc)
	for ok1 && ok2 {
		fmt.Println("Comparing : ", val1, val2)
		if val1 < val2 {
			arrayc <- val1
			val1, ok1 = <-left
		} else {
			arrayc <- val2
			val2, ok2 = <-right
		}
	}

	for ok1 {
		arrayc <- val1
		val1, ok1 = <-left
	}
	for ok2 {
		arrayc <- val2
		val2, ok2 = <-right
	}
}

func MergeSort(array []int, arrayc chan int) {
	fmt.Println("Array : ", array)
	if len(array) < 2 {
		arrayc <- array[0]
		defer close(arrayc)
		return
	}

	mid := len(array) / 2
	left := make(chan int)
	right := make(chan int)
	go MergeSort(array[:mid], left)
	go MergeSort(array[mid:], right)
	MergeChannel(left, right, arrayc)
}

func main() {
	array := []int{6, 2, 9, 1, 7, 4, 8, 3, 9, 5}
	arrayc := make(chan int)
	fmt.Println("Array Before Sorting : ", array)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		go MergeSort(array, arrayc)
	}()

	wg.Wait()

	var final []int
	for val := range arrayc {
		final = append(final, val)
	}
	fmt.Println("Array After Sorting : ", final)
}
