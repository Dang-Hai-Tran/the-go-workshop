package main

import (
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, er := f.WriteString("Test this program.")
	if er != nil {
		panic(err)
	}
}