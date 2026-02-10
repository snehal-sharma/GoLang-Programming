package main

import (
	"fmt"
	"sync"
)

func main() {
	// Channels for signaling between odd and even goroutines
	oddChan := make(chan struct{}) // Use struct{} for empty signal
	evenChan := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)
	const limit = 10

	// Odd number goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= limit; i += 2 {
			<-oddChan // Wait for signal to print odd
			fmt.Println("Odd:", i)
			if i+1 <= limit {
				evenChan <- struct{}{} // Signal even to print
			}
		}
		close(evenChan)

	}()

	// Even number goroutine
	go func() {
		defer wg.Done()
		// Range over the even channel to receive signals until it is closed
		for i := 2; i <= limit; i += 2 {
			<-evenChan // Wait for signal to print even
			fmt.Println("Even:", i)
			if i+1 <= limit {
				oddChan <- struct{}{} // Signal odd to print
			}
		}
		close(oddChan)

	}()

	oddChan <- struct{}{}

	// Wait for both goroutines to finish
	wg.Wait()

}
