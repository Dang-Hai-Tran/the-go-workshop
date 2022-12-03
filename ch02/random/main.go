package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main()  {
	var rNum int
	rand.Seed(time.Now().UnixMilli())
	for rNum = 1; rNum % 3 != 0; rNum = rand.Intn(100){
		fmt.Println(rNum)	
	}
	fmt.Println("Stop", rNum)
}