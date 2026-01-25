package main

import "fmt"

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivot := Partition(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := 0
	for j := 0; j < high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func main() {
	arrayN := []int{8, 2, 6, 3, 5, 7, 1, 4, 9}
	fmt.Println("Original Array : ", arrayN)
	QuickSort(arrayN, 0, len(arrayN)-1)
	fmt.Println("Sorted Array : ", arrayN)
}
