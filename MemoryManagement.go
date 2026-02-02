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

/*
go build -gcflags="-m -l" main.go
stackAllocation: The variable x is created, used, and its value is returned within the function's scope. The compiler can guarantee that x's lifetime does not exceed the function's execution, so it allocates x on the stack, which is fast and automatically cleaned up when the function returns.
heapAllocation: The function returns a pointer to x, meaning the variable can be accessed by code outside of the function (in this case, main). Because x must "outlive" the function's stack frame, the compiler moves x to the heap, where the Go garbage collector manages its memory. 
*/
