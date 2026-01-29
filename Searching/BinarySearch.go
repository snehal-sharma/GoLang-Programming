
package main

import "fmt"

func BinarySearchIterative(arr []int, low, high, data int) int {

	for low <= high {
		mid := low + (high-low)/2
		fmt.Printf("\nLow : %d Mid : %d High : %d", low, mid, high)
		if arr[mid] == data {
			return mid
		} else if arr[mid] < data {
			fmt.Println("\nArray : ", arr[mid+1:])
			low = mid + 1
		} else {
			fmt.Println("\nArray : ", arr[low:mid])
			high = mid - 1
		}
	}
	return -1
}

func BinarySearchRecursive(arr []int, low, high, data int) int {
	mid := low + (high-low)/2
	fmt.Printf("\nLow : %d Mid : %d High : %d", low, mid, high)

	if low > high {
		fmt.Println("Returning -1")
		return -1
	}
	if arr[mid] == data {
		return mid
	} else if arr[mid] < data {
		fmt.Println("\nArray : ", arr[mid+1:])
		return BinarySearchRecursive(arr, mid+1, high, data)
	} else {
		fmt.Println("\nArray : ", arr[low:mid])
		return BinarySearchRecursive(arr, low, mid-1, data)
	}

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	found := BinarySearchIterative(array, 0, len(array)-1, 9)
	//found := BinarySearchRecursive(array, 0, len(array)-1, 9)
	fmt.Println("\nFound Item at position : ", found)
}
