package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	maxWorkers := 3
	sem := semaphore.NewWeighted(int64(maxWorkers))
	ctx := context.Background()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := sem.Acquire(ctx, 1); err != nil {
				fmt.Println("Failed to acquire semaphore lock")
				return
			}
			defer sem.Release(1)

			fmt.Printf("\nTask %d Started ", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("\nTask %d Ended ", id)
		}(i)
	}

	wg.Wait()

	fmt.Println("All goroutines are finished")
}
