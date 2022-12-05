package main

import "fmt"

func main() {
	arr := []string{"One", "Two", "Three", "Four", "Five"}
	index := 2
	newArr := append(arr[:index], arr[index + 1:]...)
	fmt.Printf("%#v\n", newArr)
}