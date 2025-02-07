
package main

import "fmt"

func main() {

	arr := [7]int{1199, 35, -22, 1002, 57, 88, 1003}

	fmt.Println(arr)
	SecondLargest(arr[:])

}

func SecondLargest(arr []int) {
	var max, prev int
	max = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			prev = max
			max = arr[i]
		} else if arr[i] > prev && arr[i] < max {
			prev = arr[i]
		}
	}
	fmt.Printf("Max : %d \t Second Largest : %d", max, prev)

}
