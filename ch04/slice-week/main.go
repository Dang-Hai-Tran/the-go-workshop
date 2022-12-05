package main

import "fmt"

func main() {
	arr := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	newarr := arr[1:(len(arr) - 1)]
	newarr = append([]string{"Sunday"}, newarr...)
	fmt.Println(newarr)
}