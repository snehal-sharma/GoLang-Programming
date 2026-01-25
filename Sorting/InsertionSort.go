package main

import "fmt"

func InsertionSort(arr []int) {
	var i, j, value int

	for i = 1; i < len(arr); i++ {
		value = arr[i]
		j = i
		for j >= 1 && arr[j-1] > value {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = value
	}
}

func main() {
	array := []int{7, 2, 6, 3, 8, 1, 5, 9, 4}
	fmt.Println("Array Before Sorting : ", array)
	InsertionSort(array[:])
	fmt.Println("\nArray After Sorting : ", array)
}
