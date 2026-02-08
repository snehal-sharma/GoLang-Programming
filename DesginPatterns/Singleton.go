// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

type SingleInstance struct {
	Key   string
	Value int
}

var single *SingleInstance

var mutex sync.Mutex

func getSingleInstance(wg *sync.WaitGroup) *SingleInstance {
	defer wg.Done()
	if single == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if single == nil {
			fmt.Println("Creating Single Instance")
			single = &SingleInstance{Key: "aaa", Value: 20}
		} else {
			fmt.Println("Single Instance already created")
		}
	}
	return single
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getSingleInstance(&wg)
	}

	wg.Wait()

}
