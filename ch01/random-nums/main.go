package main

import (
	"fmt"
	"hash/maphash"
	"math/rand"
)

func generateRandom(a, b int) int {
	seed := new(maphash.Hash).Sum64()
	rand.Seed(int64(seed))
	return a + rand.Intn(b-a)
}

func main() {
	arr := []int{}
	for i := 1; i <= 10; i++ {
		arr = append(arr, generateRandom(10, 100))
	}
	fmt.Println(arr)
}
