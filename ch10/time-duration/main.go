package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	args := os.Args
	start := time.Now()
	cmd := exec.Command(args[1], args[2:]...)
	stdout, err := cmd.Output()
	end := time.Now()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
	elapsed := end.Sub(start).Seconds()
	fmt.Println("The program start at:", start)
	fmt.Println("The program end at:", end)
	fmt.Printf("Execution time = %f(s)\n", elapsed)
}