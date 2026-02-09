```
ctx1 := context.Background()
ctx2, cancel := context.WithCancel(ctx1)
ctx3, _ := context.WithTimeout(ctx2, time.Second)
```
* Internally, WithCancel creates a *cancelCtx.
```
type cancelCtx struct {
    Context            // parent
    done chan struct{} // read only channel closed on cancel
    err  error
    mu   sync.Mutex
    children map[canceler]struct{}
}

```
* When cancellation is  triggered
```
func (c *cancelCtx) cancel(err error) {
    c.mu.Lock()
    if c.err != nil {
        c.mu.Unlock()
        return // already canceled
    }

    c.err = err
    close(c.done) // 🔥 broadcast signal

    for child := range c.children {
        child.cancel(err)
    }
    c.children = nil
    c.mu.Unlock()
}

```
**Creation**:
* When you call a function like context.WithCancel(parent), a new cancelCtx is created, linked to its parent context, and registered as a child of the parent. It also returns a CancelFunc.

**Propagation Setup:**
* If the parent is also a cancelable context, the new context is added to the parent's children map.
* If the parent is not cancelable (e.g., context.Background()), a dedicated goroutine is started to wait for the parent's Done() channel to close or the child's Done() channel to close, to ensure proper propagation if the parent were to unexpectedly become canceled.

**Signaling Cancellation:**
* When the associated CancelFunc is called (or a deadline/timeout is reached), the cancel method of the cancelCtx is invoked.
* The cancel method acquires a lock, sets the context's err field to indicate the cause of cancellation, and closes the done channel. Closing a channel is a broadcast mechanism in Go; all goroutines waiting on that channel's receive operation are immediately unblocked.

**Recursive Cancellation:**
* The cancel method then iterates through all registered children and recursively calls their cancel methods. This ensures that the cancellation signal propagates down the entire context tree.

**Cleanup:**
* The context also unregisters itself from its parent's children map, allowing associated resources (like timers or the context itself) to be garbage collected, preventing memory leaks. 

