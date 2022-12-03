package main

import (
	"fmt"
	"runtime"
)

func main() {
	var list []uint8
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Total Alloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}