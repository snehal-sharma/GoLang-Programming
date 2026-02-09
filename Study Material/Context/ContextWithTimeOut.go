
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go performTask(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Task Timed out")
	}

}

func performTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Performing Task")
	}
}
