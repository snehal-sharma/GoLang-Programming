* Defer
  Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns.
  The arguments to the deferred function (which include the receiver if the function is a method) are evaluated when the defer executes, not when the call executes.
```
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```

* Output
Deferred functions are executed in LIFO order, so this code will cause 4 3 2 1 0 to be printed when the function returns.

```func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}```

* Output
```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```
