
package main

import "fmt"

func main() {

	arr := [7]int{10, 67, -22, 74, 57, 88, 1003}

	fmt.Println(arr)
	MinMax(arr[:])

}

func MinMax(arr []int) {
	var max, min int
	max = arr[0]
	min = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Printf("Max : %d \t Min : %d", max, min)

}
