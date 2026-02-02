package main

import (
	"fmt"
	"sync"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]int
}

func (s *Store) Get(key string) int {
	s.mu.RLock() // shared lock
	defer s.mu.RUnlock()
	return s.data[key]
}

func (s *Store) Set(key string, value int) {
	s.mu.Lock() // exclusive lock
	defer s.mu.Unlock()
	s.data[key] = value
}

func main() {
	sObject1 := Store{data: map[string]int{"aaa": 1}}
	fmt.Println("Object 1 : ", sObject1.Get("aaa"))
	sObject2 := Store{data: map[string]int{"ccc": 2}}
	sObject2.Set("ccc", 4)
	fmt.Println("Object 2 : ", sObject2.Get("ccc"))
}
