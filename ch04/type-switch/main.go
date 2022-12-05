package main

import (
	"errors"
	"fmt"
)

func doubler(v interface{}) (string, error) {
	switch t := v.(type) {
	case string:
		return t + t, nil
	case bool:
		if t {
			return "truetrue", nil
		} else {
			return "falsefalse", nil
		}
	case int:
		return fmt.Sprintf("%d%d", t, t), nil
	case float64, float32:
		return fmt.Sprintf("%f%f", t, t), nil
	default:
		return "", errors.New("unsupported type passed")
	}
}

func main() {
	res, _ := doubler("Hello")
	fmt.Println("Hello :", res)
	res, _ = doubler(-5)
	fmt.Println("-5 :", res)
	res, _ = doubler(3.14)
	fmt.Println("3.14 :", res)
	_, err := doubler([]int{1, 2, 3, 4})
	fmt.Println("[]int{1, 2, 3, 4} :", err)
}
