//Since we need the last occurrence of the data we check arr[mid] <= data first because if the data itself is at a[mid] we want to search in the second half of the array

package main

import "fmt"

func BinarySearchLastOccurrence(arr []int, low, high, data int) int {
	if high >= low {
		mid := low + (high-low)/2
		fmt.Printf("\nLow : %d Mid : %d High : %d", low, mid, high)

		if (mid == high && arr[mid] == data) || (arr[mid+1] > data && arr[mid] == data) {
			return mid
		} else if arr[mid] <= data {
			fmt.Println("\nArray : ", arr[mid+1:])
			return BinarySearchLastOccurrence(arr, mid+1, high, data)
		} else {
			fmt.Println("\nArray : ", arr[low:mid])
			return BinarySearchLastOccurrence(arr, low, mid-1, data)
		}
	}
	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 7, 7, 8, 9}
	found := BinarySearchLastOccurrence(array, 0, len(array)-1, 7)
	fmt.Println("\nFound last occurrence at : ", found)
}
