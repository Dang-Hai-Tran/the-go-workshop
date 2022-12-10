package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("person.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	fmt.Println("Read line by line")
	for {
		// Read line by line
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}
