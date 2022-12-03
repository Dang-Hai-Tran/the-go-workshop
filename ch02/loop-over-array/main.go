package main

import "fmt"

func main() {
	name := []string{"Jim", "Jane", "Joe", "June"}
	for _, v := range name {
		fmt.Println(v)
	}
}