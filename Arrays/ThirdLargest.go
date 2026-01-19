
package main

import "fmt"

func main() {

	arr := [7]int{11, 35, -22, 1002, 57, 88, 1003}

	fmt.Println(arr)
	ThirdLargest(arr[:])

}

func ThirdLargest(arr []int) {
	var first, second, third int
	first = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > first {
			third = second
			second = first
			first = arr[i]
		} else if arr[i] > second {
			third = second
			second = arr[i]
		} else if arr[i] > third {
			third = arr[i]
		}
	}
	fmt.Printf("Max : %d \t Second Largest : %d \t Third Largest :%d ", first, second, third)

}
