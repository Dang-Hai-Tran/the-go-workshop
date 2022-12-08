package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	message := []byte("Nice to meet you!")
	err := ioutil.WriteFile("test.txt", message, 0744)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success.")
}