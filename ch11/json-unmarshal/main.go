package main
import (
	"fmt"
	"encoding/json"
)

type greeting struct {
	Message string
}

func main() {
	data := `{
		"message" : "Hi, nice to meet you"
	}`
	// Unmarshal to a struct
	var v greeting
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	// Unmarshal to a map
	var m map[string]string
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m["message"])
}