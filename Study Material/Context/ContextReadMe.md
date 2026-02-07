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
* context.Context is passed by value, but it contains references to shared internal state like a done channel. Canceling closes that channel, which all copies observe.
```
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
```
