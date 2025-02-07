package main

import "fmt"

func main() {

	arr := [7]int{4, 67, 22, 74, 57, 88, 98}

	fmt.Println(arr)
	found := search(arr[:], 74)
	if found != -1 {
		fmt.Printf("Value found at position %d\n", found)
	} else {
		fmt.Println("Value not found")
	}
	//Insert at position 4 Value 1001
	insert(arr[:], 1001, 4)

	//Delete at position 3
	deletePos(arr[:], 3)

}

func search(arr []int, elem int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return i
		}
	}
	return -1
}

func insert(arr []int, elem int, pos int) {
	arr = append(arr, 0)
	fmt.Println(arr, len(arr))
	for i := len(arr) - 1; i > pos; i-- {
		arr[i] = arr[i-1]
	}
	arr[pos] = elem
	fmt.Println(arr, len(arr))
}

func deletePos(arr []int, pos int) {
	fmt.Println(arr, len(arr))
	for i := pos; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr = arr[:len(arr)-1]
	fmt.Println(arr, len(arr))
}
