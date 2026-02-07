* Context is immutable
* context.WithValue returns a new context. You can’t add values to an existing context. Each function must derive a new context using context.WithValue and return it.
* ctx is passed by value. In the example below the caller never recieves the new context.
```
func a(ctx context.Context) {
	ctx = context.WithValue(ctx, "x", 1)
}
```
* Always return the derived context
* Never store context in structs. Because context is request-scoped and cancelable. Storing it in structs causes lifetime mismatches, broken cancellation, hidden dependencies, and race conditions.
