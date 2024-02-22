package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println()
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		// Allocating 50,000,000 bytes
		s := make([]byte, 50000000)
		println(len(s))
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		// Allocating 100,000,000 bytes
		s := make([]byte, 100000000)
		println(len(s))

		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}
