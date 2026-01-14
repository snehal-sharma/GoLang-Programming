/*
Optimised bubble sort where the for loop breaks when no values are swapped during a pass signifying that the array is already sorted.
*/
package main

import "fmt"

func main() {
	temp := 0
	valuesSwapped := true
	arrayN := []int{2, 3, 6, 8, 1, 7, 9, 4, 5}
	fmt.Println("Original Array : ", arrayN)
	for pass := len(arrayN) - 1; pass >= 0 && valuesSwapped; pass-- {
		fmt.Println("Pass : ", pass)
		valuesSwapped = false
		for i := 0; i < pass; i++ {
			if arrayN[i] > arrayN[i+1] {
				temp = arrayN[i]
				arrayN[i] = arrayN[i+1]
				arrayN[i+1] = temp
				valuesSwapped = true
			}
		}
		fmt.Println(arrayN, valuesSwapped)
	}
}
