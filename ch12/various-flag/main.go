package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	s := flag.String("name", "", "Your name is an string (default : '')")
	i := flag.Int("age", 0, "Your age is an int (default : 0)")
	b := flag.Bool("married", false, "You are married ? true or false (default : false)")
	flag.Parse()
	// If we want name flag is required
	if *s == "" {
		fmt.Println("Name is required!")
		// Print default setting of all flags
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("Name :", *s)
	fmt.Println("Age :", *i)
	fmt.Println("Married :", *b)
}
