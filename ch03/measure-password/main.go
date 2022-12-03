package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	pwR := []rune(pw)
	if len(pwR) < 8 || len(pwR) > 15 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	pw := "ThisI$5A"
	if passwordChecker(pw) {
		fmt.Println(pw, "is good password.")
	} else {
		fmt.Println(pw, "is bad password.")
	}
}
