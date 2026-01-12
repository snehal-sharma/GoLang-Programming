//Implement GoLang Channels using WaitGroup and Channel Directions
package main

import (
	"fmt"
	"sync"
)

func computeSquare(n int, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- n * n
}

func main() {
	arrayN := []int{1, 2, 3, 4, 5}
	c := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(len(arrayN))

	for i := 0; i < len(arrayN); i++ {
		go computeSquare(arrayN[i], c, &wg)
	}

	wg.Wait()
	close(c)
	value, ok := <-c
	for ok {
		fmt.Println(value)
		value, ok = <-c
	}
}
