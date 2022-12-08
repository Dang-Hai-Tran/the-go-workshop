package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	event, err := time.Parse(time.ANSIC, "Fri Dec 2 15:03:24 2020")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if now.After(event) {
		fmt.Println("Now is after event.")
	} else {
		fmt.Println("Now is before event")
	}
}