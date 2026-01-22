package main

import "fmt"

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	fmt.Println("Array One :", array[:mid])
	arrayOne := MergeSort(array[:mid])
	fmt.Println("Array Two :", array[mid:])
	arrayTwo := MergeSort(array[mid:])
	return Merge(arrayOne, arrayTwo)

}

func Merge(arrayLeft, arrayRight []int) []int {
	arrayFinal := make([]int, 0, len(arrayLeft)+len(arrayRight))
	i, j := 0, 0

	for i < len(arrayLeft) && j < len(arrayRight) {
		if arrayLeft[i] <= arrayRight[j] {
			arrayFinal = append(arrayFinal, arrayLeft[i])
			i++
		} else {
			arrayFinal = append(arrayFinal, arrayRight[j])
			j++
		}
	}

	arrayFinal = append(arrayFinal, arrayLeft[i:]...)
	arrayFinal = append(arrayFinal, arrayRight[j:]...)

	return arrayFinal
}

func main() {
	array := []int{2, 5, 8, 6, 3, 4, 9, 1, 7}
	fmt.Println("Array Before Sorting : ", array)
	fmt.Println("Array After Sorting : ", MergeSort(array[:]))

}
