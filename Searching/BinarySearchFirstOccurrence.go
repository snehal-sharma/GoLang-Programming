
package main

import "fmt"

func BinarySearchFirstOccurrence(arr []int, low, high, data int) int {
	if low <= high {
		mid := low + (high-low)/2
		fmt.Printf("\nLow : %d Mid : %d High : %d", low, mid, high)

		if (mid == low && arr[mid] == data) || (arr[mid-1] < data && arr[mid] == data) {
			return mid
		} else if arr[mid] < data {
			fmt.Println("\nArray : ", arr[mid+1:])
			return BinarySearchFirstOccurrence(arr, mid+1, high, data)
		} else {
			fmt.Println("\nArray : ", arr[low:mid])
			return BinarySearchFirstOccurrence(arr, low, mid-1, data)
		}
	}
	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7, 8, 9}
	found := BinarySearchFirstOccurrence(array, 0, len(array)-1, 4)
	fmt.Println("\nFound first occurrence at : ", found)
}
