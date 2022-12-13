package main
import (
	"fmt"
)

func greeting(ch chan string) {
	mes := <- ch
	ch <- fmt.Sprintf("Nice to meet you %s\n", mes)
	ch <- fmt.Sprintln("Hello Dang")
}

func main() {
	ch := make(chan string)
	go greeting(ch)
	ch <- "David"
	fmt.Print(<- ch)
	fmt.Print(<- ch)
}