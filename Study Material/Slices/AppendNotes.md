The append function in Go is a built-in function used to add elements to the end of a slice. It handles the dynamic resizing of the underlying array as needed and always returns a new slice header, which must be stored in a variable, typically the original slice variable itself. 

**How append Works**
* Basics : Length = number of elements the slice currently exposes. Capacity = how far the slice can grow before reallocation.
* The behavior of append is primarily determined by the capacity of the original slice's underlying array. Slices in Go are dynamic, but they are backed by fixed-size arrays. A slice has both a length (the number of accessible elements) and a capacity (the total number of elements in the underlying array). 
When you call append(s, elements...):

**Check Capacity**: Go checks if the underlying array of slice s has enough remaining capacity (capacity - length) to accommodate the new elements.

# Sufficient Capacity (In-place append):
* If there is enough space, the new elements are copied into the existing underlying array immediately after the current length of the slice.
* A new slice header (which includes a pointer to the same underlying array, a new length, and the original capacity) is returned. The original underlying array is mutated in the previously unused memory slots. This can cause side effects if other slices share the same backing array and have a "window" into the modified part.
```
s := make([]int, 2, 5)
s = []int{1, 2}
s = append(s, 3)
```
**Internally**
* Check: len + 1 <= cap 
* Write element at index len
* Increment len
* Return updated slice header


# Insufficient Capacity (New allocation):
* If there isn't enough capacity, Go allocates a new, larger underlying array (typically double the original capacity, with some internal optimizations for growth heuristics).
* All elements from the original slice are copied to the new array, followed by the new elements.
* A new slice header pointing to this new, larger array is returned. The original array remains unchanged. 

# Key Characteristics and Usage
* Reassignment is Crucial: Because append might return a slice with a different underlying array, you must assign the result back to the original slice variable to use the updated slice correctly:

```
slice = append(slice, element)
```

# Why is it important to save slice header in the same variable.
```
package main

import "fmt"

func main() {
	slice := make([]int, 0, 5)
	slice = append(slice, 1)
	_ = append(slice, 2)
	slice = append(slice, 3)
	fmt.Println(slice, len(slice), cap(slice))

}
```
**Output**
```
[1 3] 2 5
```
* Append only modifies the backing array and returns a new slice header. At append(slice, 2) You ignored the returned header, so the slice length never changed. The element does get written into the array — but your slice still thinks its length is 1. Hence at slice = append(slice, 3) is written at len(slice)+1, essentially overwriting 2.

* **Variadic Function**: append is a variadic function, meaning it can take zero or more elements.
To append a single element: slice = append(slice, 42).
To append multiple elements: slice = append(slice, 1, 2, 3).
To append the elements of one slice to another, use the ... (spread) operator: slice1 = append(slice1, slice2...).

* Works with nil slices: You can call append on a nil slice, and Go will handle the initial memory allocation, effectively creating a new slice.
