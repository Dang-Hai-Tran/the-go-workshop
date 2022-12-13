package main

import (
	"fmt"
	"time"
)

func next(ptr *int) {
	*ptr++
	fmt.Println(*ptr)
}

func main() {
	a := 0;
	go next(&a)
	go next(&a)
	go next(&a)
	go next(&a)
	time.Sleep(10 * time.Microsecond)
	fmt.Println("Value a =", a)
}