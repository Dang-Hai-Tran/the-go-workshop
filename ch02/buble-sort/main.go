package main
import (
	"fmt"
)

func main() {
	arr := []int {5, 8, 2, 0, 1, 3, 7, 9, 6}
	for i := 0; i < len(arr) - 1;{
		if (arr[i] > arr[i+1]){
			arr[i], arr[i+1] = arr[i+1], arr[i]
			i = 0
		} else {
			i++
		}
	}
	fmt.Println(arr)
}