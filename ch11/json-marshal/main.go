package main

import (
	"fmt"
	"encoding/json"
)

type greeting struct {
	SomeMessage string
}

func main() {
	// Marshal a struct
	var v greeting
	v.SomeMessage = "Marshall me!"
	js, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", js)
	// Marshal a map
	m := map[string]int{"students": 500, "teacher": 20}
	js, err = json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", js)
}