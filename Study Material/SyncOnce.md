* sync.Once ensures that a function runs exactly once, even if multiple goroutines call Do() concurrently or some goroutines arrive while the function is still running.
**Internal Structure**
```
type Once struct {
	done uint32
	m    Mutex
}
```
* done → atomic flag (0 = not run, 1 = run)
* m → mutex for slow path synchronization

# Breakdown
**Fast Path**
* If done==1
```
if atomic.LoadUint32(&o.done) == 1 {
	return
}
```
* If done==0
**Slow Path**
```
o.m.Lock()
defer o.m.Unlock()

if o.done == 0 {
	f()
	atomic.StoreUint32(&o.done, 1)
}

```
* If f() panics done is NOT set. Next call to Do() will try again.


