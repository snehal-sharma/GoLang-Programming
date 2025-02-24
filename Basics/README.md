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
