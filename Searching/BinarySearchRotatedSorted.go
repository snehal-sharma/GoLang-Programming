package main

import "fmt"

func BinarySearchRotatedRecursive(array []int, low, high, data int) int {
	mid := low + (high-low)/2

	if array[mid] == data {
		return mid
	}

	if array[low] <= array[mid] {
		fmt.Println("Array Iteration: ", array[low:mid+1])
		if array[mid] > data && array[low] <= data {
			fmt.Println("Array :", array[low:mid])
			return BinarySearchRotatedRecursive(array, low, mid-1, data)
		} else {
			fmt.Println("Array :", array[mid+1:high+1])
			return BinarySearchRotatedRecursive(array, mid+1, high, data)
		}
	} else {
		fmt.Println("Array Iteration: ", array[mid:high+1])
		if array[mid] < data && array[high] >= data {
			fmt.Println("Array :", array[mid+1:high+1])
			return BinarySearchRotatedRecursive(array, mid+1, high, data)
		} else {
			fmt.Println("Array :", array[low:mid])
			return BinarySearchRotatedRecursive(array, low, mid-1, data)
		}
	}
}

func BinarySearchRotated(array []int, data int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)/2
		if array[mid] == data {
			return mid
		}

		if array[low] <= array[mid] {
			fmt.Println("Array : ", array[low:mid+1])
			if array[mid] > data && array[low] <= data {
				fmt.Println("Array :", array[low:mid])
				high = mid - 1
			} else {
				fmt.Println("Array :", array[mid+1:high+1])
				low = mid + 1
			}
		} else {
			fmt.Println("Array : ", array[mid:high+1])
			if array[mid] < data && array[high] >= data {
				fmt.Println("Array : ", array[mid+1:high+1])
				low = mid + 1
			} else {
				fmt.Println("Array :", array[low:mid])
				high = mid - 1
			}
		}
	}
	return -1
}

func main() {
	arrayN := []int{6, 7, 8, 1, 2, 3, 4, 5}
	//position := BinarySearchRotatedRecursive(arrayN, 0, len(arrayN)-1, 7)
	position := BinarySearchRotated(arrayN, 7)
	fmt.Println("Found Item at Position : ", position)

}
