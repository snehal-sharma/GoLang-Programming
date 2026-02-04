In Go, everything is passed by value, including slices. However, a slice's value is a header that contains a pointer to an underlying array, its length, and its capacity. 

```
type sliceHeader struct {
    Data uintptr // pointer to the underlying array
    Len  int     // number of elements in the slice
    Cap  int     // capacity of the slice
}
```


This leads to a behavior that is often described as "pass by reference" with respect to the elements in the slice: 

* **Modifying existing elements**: If you modify the elements of a slice inside a function (e.g., s[0] = 10), the original slice will see the changes because both the original and the copied slice headers point to the same underlying array.
* **Modifying the slice header (length/capacity)**: If you use a function like append() that changes the slice's length or capacity, the change will not be reflected in the original slice. This is because you are modifying the copy of the slice header within the function's scope, not the original. 

**Summary**
* To effectively work with slices in functions:
* **To modify existing elements:** Pass the slice by value (the default). The changes to the underlying array elements will be visible to the caller.
* **To modify the slice's structure (length or capacity):**
* Return the new slice from the function and assign it back to the original variable in the caller function (e.g., mySlice = modify(mySlice)).
* Alternatively, pass a pointer to the slice to the function if you need to modify the slice's header in place.

