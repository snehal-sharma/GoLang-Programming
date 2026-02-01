package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var random int
	arrayN := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := len(arrayN) - 1; i >= 0; i-- {
		if i == 0 {
			i = 1
			random = rand.Intn(i)
			i--
		} else {
			random = rand.Intn(i)
		}
		arrayN[i], arrayN[random] = arrayN[random], arrayN[i]
	}
	fmt.Println("Array : ", arrayN)
}
