/*
We use *SingleInstance because a singleton must represent exactly one shared instance. 
Returning a pointer avoids copying, preserves shared state across goroutines, and is the idiomatic and safe way to implement a singleton in Go.
*/
package main

import (
	"fmt"
	"sync"
)

type SingleInstance struct {
	Key   string
	Value int
}

var (
	single *SingleInstance
	once   sync.Once
)

func getSingleInstance() *SingleInstance {
	once.Do(func() {
		fmt.Println("Creating single instance")
		single = &SingleInstance{Key: "aaa", Value: 20}
	})
	return single
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = getSingleInstance()
		}()
	}

	wg.Wait()
}
