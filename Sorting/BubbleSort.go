/*
Optimised bubble sort where the for loop breaks when no values are swapped during a pass signifying that the array is already sorted.
*/
package main

import "fmt"

func BubbleSort(array []int) {
	temp := 0
	valuesSwapped := true
	for pass := len(array) - 1; pass >= 0 && valuesSwapped; pass-- {
		fmt.Println("Pass : ", pass)
		valuesSwapped = false
		for i := 0; i < pass; i++ {
			if array[i] > array[i+1] {
				temp = array[i]
				array[i] = array[i+1]
				array[i+1] = temp
				valuesSwapped = true
			}
		}
		fmt.Println(array, valuesSwapped)
	}

}

func main() {
	array := []int{2, 3, 6, 8, 1, 7, 9, 4, 5}
	fmt.Println("Array Before Sorting : ", array)
	BubbleSort(array[:])
	fmt.Println("\nArray After Sorting : ", array)
}
