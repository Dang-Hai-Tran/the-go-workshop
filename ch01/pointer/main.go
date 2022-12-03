package main
import (
	"fmt"
)

func swap(a *int, b *int) {
	*a, *b = *b, *a	
}

func main() {
	var a, b = 5, 10
	swap(&a, &b)
	fmt.Printf("a = %d, b= %d", a, b)
}