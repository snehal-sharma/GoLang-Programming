package main

import "fmt"

func main() {

	arr := [6]int{11, 35, 1002, 57, 88, 1003}
	Reverse(arr[:])
	fmt.Println(arr)

}

func Reverse(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		temp := arr[left]
		arr[left] = arr[right]
		arr[right] = temp
		left++
		right--
	}

}
