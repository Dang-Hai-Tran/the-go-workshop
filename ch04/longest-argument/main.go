package main

import (
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {
	args := os.Args
	// Check number arguments greater or equal than minArgs
	if (len(args) < minArgs) {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}
	return args
}

func findLongestArg(args []string) string {
	if len(args) == 1 {
		fmt.Println("There aren't any argument")
		os.Exit(1)
	}
	var longest string = args[1]
	for i:=1; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

func main() {
	args := getPassedArgs(3)
	longest := findLongestArg(args)
	fmt.Println("The longest argument is", longest)
}