* A goroutine is a lightweight concurrent execution unit managed by the Go runtime Scheduler. 
* The stacks start small with an initial stack size (usually 2KB or 4KB) so they are cheap and grow by allocating (and freeing) heap storage as required.
  As a result, Go programs can spawn thousands of goroutines with relatively low overhead.
* Goroutines are multiplexed onto smaller number of OS threads in an interleaved manner so if one should block, such as while waiting for I/O, others continue to run.

**GMP Model**
* https://leapcell.io/blog/unveiling-go-s-scheduler-secrets-the-g-m-p-model-in-action

* Go can handle hundreds of thousands or even millions of goroutines, as they are lightweight and managed efficiently by the Go runtime. However, the practical limit is determined by the available system resources, primarily memory and CPU, and the nature of the tasks the goroutines are performing. 

* An unbuffered channel in Go is a synchronization point with zero capacity that requires both a sender and a receiver to be ready at the same time for communication to succeed. UseCase:-Signaling between goroutines,Ensuring one step finishes before the next,Worker start/stop signals.
* Sending to an unbuffered channel without recieving it causes a goroutine leak and subsequent deadlock. Receiving from a channel without sending anything in returns a zero value. Sending/Receiving on a nil channel blocks forever(deadlock).

* A buffered channel in Go is a channel with a fixed capacity to store a predetermined number of values without requiring an immediate corresponding receive operation. This allows for a degree of asynchronous communication between goroutines, decoupling the sender and receiver's timing. UserCase:-Worker pools, Job queues, Rate limiting, Logging pipelines.
* Sending to a buffered channel when it is full causes a deadlock. Receiving from an empty channel returns 0.


