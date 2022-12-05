package main

import "fmt"

func geneSlices() ([]int, []int, []int) {
	var s1 []int
	// Use make help us to control the capacity of hidden array
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)
	return s1, s2, s3
}

func main() {
	s1, s2, s3 := geneSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
} 