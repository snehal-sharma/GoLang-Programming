package main

import (
	"fmt"
	"time"
)

func main() {
	tokens := make(chan struct{}, 5)
	const (
		rate  = 2
		burst = 5
	)

	for i := 0; i < burst; i++ {
		tokens <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(time.Second / rate)
		defer ticker.Stop()

		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default:
			}
		}
	}()

	for i := 0; i <= 10; i++ {
		<-tokens
		fmt.Printf("\nRequest %d processed at %v", i, time.Now())
	}
}
