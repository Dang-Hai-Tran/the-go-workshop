package main
import (
	"fmt"
)

func printNums(nums... int)  {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main()  {
	printNums(1, 2, 3, 4, 5, 6)
}