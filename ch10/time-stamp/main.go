package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Format(time.ANSIC)
	fmt.Println(now)
}