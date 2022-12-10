package parsedata

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	ID int
	Payee string
	Spent float64
	Category BudgetCategory
}

type BudgetCategory string
const  (
	AutoFuel BudgetCategory = "autoFuel"
	Food BudgetCategory = "food"
	Mortgage BudgetCategory = "mortgage"
	Repairs	BudgetCategory = "repairs"
	Insurance BudgetCategory = "insurance"
	Utilities BudgetCategory = "utilities"
)
var ErrInvalidTransactionCategory error = errors.New("transaction category not found")

func convertTransactionCategoryToBudgetCategory(transactionCategory string) (BudgetCategory, error) {
	switch transactionCategory {
	case "fuel", "gas":
		return AutoFuel, nil
	case "food":
		return Food, nil
	case "mortgage":
		return Mortgage, nil
	case "repairs":
		return Repairs, nil
	case "car insurance", "life insurance", "retirement":
		return Insurance, nil
	case "utilities", "tv":
		return Utilities, nil
	default:
		return "", ErrInvalidTransactionCategory
	}
}

func ReadTransactionFile(file *os.File) ([]Transaction, error) {
	r := csv.NewReader(file)
	header := true
	listTransaction := []Transaction{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []Transaction{}, err
		}
		if !header {
			var newTransaction Transaction
			for i, v := range record {
				v = strings.TrimSpace(v)
				switch i {
				case 0:
					newTransaction.ID, err = strconv.Atoi(v)
					if err != nil {
						return []Transaction{}, err
					}
				case 1:
					newTransaction.Payee = v
				case 2:
					newTransaction.Spent, err = strconv.ParseFloat(v, 64)
					if err != nil {
						return []Transaction{}, err
					}
				case 3:
					newTransaction.Category, err = convertTransactionCategoryToBudgetCategory(v)
					if err != nil {
						return []Transaction{}, err
					}
				}

			}
			listTransaction = append(listTransaction, newTransaction)
		}
		header = false
	}
	return listTransaction, nil
}