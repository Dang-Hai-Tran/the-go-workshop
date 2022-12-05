package main
import (
	"fmt"
)

type calcul func(int, int) string

func add(i, j int) string {
	return fmt.Sprintf("Add %d and %d is %d", i, j, i + j)
}

func sub(i, j int) string {
	return fmt.Sprintf("Subtract %d and %d is %d", i, j, i - j)
}

func calculator(f calcul, i int, j int) string {
	return f(i, j)
}

func main() {
	fmt.Println(calculator(add, 10, 5))
	fmt.Println(calculator(sub, 10, 5))
}