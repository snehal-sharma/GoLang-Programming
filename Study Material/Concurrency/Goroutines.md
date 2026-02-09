* A goroutine is a lightweight concurrent execution unit managed by the Go runtime Scheduler. 
* The stacks start small with an initial stack size (usually 2KB or 4KB) so they are cheap and grow by allocating (and freeing) heap storage as required.
  As a result, Go programs can spawn thousands of goroutines with relatively low overhead.
* Goroutines are multiplexed onto smaller number of OS threads in an interleaved manner so if one should block, such as while waiting for I/O, others continue to run.
**GMP Model**
* https://leapcell.io/blog/unveiling-go-s-scheduler-secrets-the-g-m-p-model-in-action
