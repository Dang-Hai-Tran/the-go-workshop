package main

import (
	"fmt"
	"os"
)

func main() {
	file := "newText.txt"
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Printf("File %v was created successfully.\n", file)
	text := "Hello World\n"
	f.Write([]byte(text))
}
