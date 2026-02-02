package main

import "fmt"

//go:noinline
func stackAllocation() int {
	x := 42 // Allocated on the stack.
	return x // Return the value (not a reference).
}

//go:noinline
func heapAllocation() *int {
	x := 42 // Initially on stack.
	return &x // Return the address - x escapes to the heap.
}

func main() {
	val := stackAllocation()
	fmt.Printf("Stack allocation value: %d\n", val)

	ptr := heapAllocation()
	fmt.Printf("Heap allocation value: %d\n", *ptr)
}
