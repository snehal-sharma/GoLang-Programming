package main

import "fmt"

func SelectionSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		temp := array[i]
		array[i] = array[min]
		array[min] = temp

		fmt.Println("Pass : ", i)
		fmt.Println(array)
	}

}

func main() {
	array := []int{3, 6, 8, 1, 7, 4, 9, 2, 5}
	fmt.Println("Array Before Sorting : ", array)
	SelectionSort(array[:])
	fmt.Println("\nArray After Sorting : ", array)
}
