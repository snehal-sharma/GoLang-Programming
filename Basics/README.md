# GoLang Basics

* Package Imports in Go
  ```
  import (
	"fmt"
	"math/rand"
  )
  ```
* In Go, a name is exported if it begins with a capital letter. When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
* Go's return values may be named. If so, they are treated as variables defined at the top of the function. A return statement without arguments returns the named return values. This is known as a "naked" return.
  ```
  func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
  }
  ```
* When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
  ```
  i := 42           // int
  f := 3.142        // float64
  g := 0.867 + 0.5i // complex128
  ```
* Like for, the if statement can start with a short statement to execute before the condition.
  ```
  func pow(x, n, lim float64) float64 {
  	if v := math.Pow(x, n); v < lim {
  		return v
  	} else {
  		fmt.Printf("%g >= %g\n", v, lim)
  	}
  	// can't use v here, though
  	return lim
  }
  ```
* A defer statement defers the execution of a function until the surrounding function returns. Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
  ```
  func main() {
  	fmt.Println("counting")
  
  	for i := 0; i < 10; i++ {
  		defer fmt.Println(i)
  	}
  
  	fmt.Println("done")
  }
  ```
  Output
  ```
  counting
  done
  9
  8
  7
  6
  5
  4
  3
  2
  1
  0
  ```
* Array : An array's length is part of its type, so arrays cannot be resized. An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
  ```
  var a [2]string
  primes := [6]int{2, 3, 5, 7, 11, 13}
  ```
* Changing the elements of a slice modifies the corresponding elements of its underlying array.
* A slice has both a length and a capacity. The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
  Ref: https://go.dev/tour/moretypes/11
* Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
  ```
  a := make([]int, 5)
  ```
* Append : It is common to append new elements to a slice, and so Go provides a built-in append function.
  ```
  func append(s []T, vs ...T) []T
  var s []int  // Nil Slice
  s = append(s, 13) // Appending to a Nil Array
  a = append(a,make([]int,5)...) // alternative resizing
  ```
* You can declare a method on non-struct types, too.
* You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package.
* https://go.dev/tour/methods/4
* You can declare methods with pointer receivers. This means the receiver type has the literal syntax *T for some type T. Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers. Example : https://go.dev/tour/methods/4
* Methods and pointer indirection
  1. Calling with a value (type T): If you call a method with a pointer receiver on a value v of type T (e.g., v.Scale(2)), Go automatically takes the address of v and interprets the call as (&v).Scale(2). This implicit conversion requires that the variable v is addressable (e.g., a variable, not a map value or an interface value's data).
    2. Calling with a pointer (&type): If you call the method with a pointer p of type *T (e.g., p.Scale(3)), it works as expected because the receiver type matches the variable type
