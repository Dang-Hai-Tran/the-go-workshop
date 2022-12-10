package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"the-go-workshop/ch12/parsing-bank-transaction/parseData"
)

type Payment struct {
	AutoFuel  []float64
	Food      []float64
	Mortgage  []float64
	Repair    []float64
	Insurance []float64
	Utility   []float64
}

func getSum(slice ...float64) float64 {
	sum := 0.0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func returnPaymentByCategory(category string) float64 {
	f, err := os.Open("./data/transaction.csv")
	if err != nil {
		panic(err)
	}
	listTransaction, err := parsedata.ReadTransactionFile(f)
	if err != nil {
		panic(err)
	}
	payment := Payment{}
	for _, transaction := range listTransaction {
		switch transaction.Category {
		case parsedata.AutoFuel:
			payment.AutoFuel = append(payment.AutoFuel, transaction.Spent)
		case parsedata.Food:
			payment.Food = append(payment.Food, transaction.Spent)
		case parsedata.Mortgage:
			payment.Mortgage = append(payment.Mortgage, transaction.Spent)
		case parsedata.Repairs:
			payment.Repair = append(payment.Repair, transaction.Spent)
		case parsedata.Insurance:
			payment.Insurance = append(payment.Insurance, transaction.Spent)
		case parsedata.Utilities:
			payment.Utility = append(payment.Utility, transaction.Spent)
		}
	}
	var res float64
	switch category {
	case "autofuel":
		res = getSum(payment.AutoFuel...)
	case "food":
		res = getSum(payment.Food...)
	case "mortgage":
		res = getSum(payment.Mortgage...)
	case "repair":
		res = getSum(payment.Repair...)
	case "insurance":
		res = getSum(payment.Insurance...)
	case "utility":
		res = getSum(payment.Utility...)
	default:
		err := errors.New("category not found")
		panic(err)
	}
	return res
}

func main() {
	c := flag.String("category", "", "This category that you want to see the total payment")
	flag.Parse()
	if *c == "" {
		fmt.Println("Flag not found")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("The payment of %v is %v\n", *c, returnPaymentByCategory(*c))
}
