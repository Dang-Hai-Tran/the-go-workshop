package main

import (
	"fmt"
	"os"
)

func checkArgs() []string {
	args := os.Args
	if (len(args) == 1) {
		fmt.Println("No argument")
		os.Exit(1)
	} else if (len(args) >= 3) {
		fmt.Println("Only one argument allowed")
		os.Exit(1)
	}
	return args
}

func findName(input string, tableNames map[string]string) string {
	var res string
	for key, value := range tableNames {
		if key == input {
			res = value
		}
	} 
	return res
}

func main() {
	tableNames := map[string]string{
		"305" : "Sue",
		"204" : "Bob",
		"631" : "Jake",
		"073" : "Tracy",
	}
	args := checkArgs()
	name := findName(args[1], tableNames)
	if name != "" {
		fmt.Println("Hi", name)
	} else {
		fmt.Println("Name not found")
		os.Exit(1)
	}
}

