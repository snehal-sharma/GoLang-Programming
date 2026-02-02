// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime"
)

func PrintMemStats(mem *runtime.MemStats, stage string) {
	runtime.ReadMemStats(mem)
	fmt.Println("Stage : ", stage)
	fmt.Printf("\nAllocation on heap :%d KB", mem.HeapAlloc/1024)
	fmt.Printf("\nTotal Allocation :%d KB", mem.TotalAlloc/1024)
	fmt.Printf("\nMemory Allocated on the system :%d KB", mem.Sys/1024)
	fmt.Println("\nNumber of garbage collection (NumGC)", mem.NumGC)
}

func main() {
	var mem runtime.MemStats

	PrintMemStats(&mem, "Before Allocation")

	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i)
	}

	PrintMemStats(&mem, "After Allocation")

	data = nil
	runtime.GC()

	PrintMemStats(&mem, "After Garbage Collection")

}

/*
runtime.MemStats struct: This struct holds various memory statistics.
runtime.ReadMemStats(&mem): This function populates the mem variable with the current memory data from the Go runtime.
mem.HeapAlloc: This field indicates the bytes of allocated heap objects currently in use.
mem.TotalAlloc: This field shows the cumulative bytes allocated by the program since it started.
mem.Sys: This indicates the total bytes of memory obtained from the operating system.
mem.NumGC: This field shows the number of completed garbage collection cycles.
runtime.GC(): Calling this function explicitly forces the garbage collector to run, reclaiming unused memory and demonstrating its effect on the HeapAlloc and Sys values. 
*/
