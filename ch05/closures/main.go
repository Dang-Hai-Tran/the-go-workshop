package main

import (
	"fmt"
)

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	increase := incrementor()
	fmt.Println(increase())
	fmt.Println(increase())
	fmt.Println(increase())
	fmt.Println(increase())
}
