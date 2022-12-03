package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(5) + 1
	str := strings.Repeat("*", randNum)
	fmt.Println(str)
}