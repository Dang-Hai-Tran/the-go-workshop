package main

import (
	"errors"
	"fmt"
)

func message(msg string) {
	if (msg == "good-bye") {
		panic(errors.New("something went wrong"))
	}
}

func main() {
	msg := "good-bye"
	defer func() {
		fmt.Println("This defer statement is executed")
	}()
	message(msg)
	// All statement apter panic isn't executed
	fmt.Println("This line isn't printed")
}