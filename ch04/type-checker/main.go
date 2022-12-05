package main

import (
	"fmt"
)

func getType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case bool:
		return "bool"
	case string:
		return "string"
	case float32, float64:
		return "float"
	default:
		return "unknown"
	}
}

func main() {
	res := getType(12314)
	fmt.Println("12314 :", res)
	res = getType("hello")
	fmt.Println("hello: ", res)
	res = getType(true)
	fmt.Println("true :", res)
	res = getType(3.124242)
	fmt.Println("3.124242 :", res)
	res = getType([]int{})
	fmt.Println("[]int{} :", res)
}
