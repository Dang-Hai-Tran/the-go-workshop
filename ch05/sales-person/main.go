package main

import (
	"fmt"
)

func itemsSold() {
	items := map[string]int{}
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24
	for name, item := range items {
		var rep string
		if (item < 35) {
			rep = "below expectation"
		} else if (item < 70) {
			rep = "meet expectation"
		} else {
			rep = "exceeded expectation"
		}
		fmt.Printf("%s sold %d items and %s.\n", name, item, rep)
	}
}

func main() {
	itemsSold()
}