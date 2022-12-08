package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("The script start at:", start)
	time.Sleep(5 * time.Second)
	end := time.Now()
	fmt.Println("The script has completed at:", end)
}