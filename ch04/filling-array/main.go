package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	for i := range arr {
		arr[i] = i + 1
	}
	fmt.Println(arr)
}
