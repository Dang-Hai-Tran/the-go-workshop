package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var creditScore, inCome, loanAmount, loanTerm int
	fmt.Printf("Credit Score : ")
	fmt.Scanln(&creditScore)
	fmt.Printf("Income : ")
	fmt.Scanln(&inCome)
	fmt.Printf("Loan Amount : ")
	fmt.Scanln(&loanAmount)
	fmt.Printf("Loan Term : ")
	fmt.Scanln(&loanTerm)
	if (creditScore <= 0 || inCome <= 0 || loanAmount <= 0 || loanTerm <= 0 || loanTerm % 12 != 0) {
		err := errors.New("INPUT NOT CORRECT")
		log.Fatal(err)
	}
	var rate float32
	if creditScore >= 450 {
		rate = 15
	} else {
		rate = 20
	}
	var monthlyPayment float32
	var approved bool
	monthlyPayment = float32(loanAmount) * (1 + rate / 100) / 12
	if creditScore >= 450 {
		if monthlyPayment <= float32(inCome) * 0.2 {
			approved = true
		} else {
			approved = false
		}
	} else {
		if monthlyPayment <= float32(inCome) * 0.1 {
			approved = true
		} else {
			approved = false
		}
	}
	fmt.Println("Approved :", approved)
}
