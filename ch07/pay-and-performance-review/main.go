package main

import (
	"fmt"
)

type T interface{}

type Employee struct {
	Id        string
	FirstName string
	LastName  string
}

type Developer struct {
	Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]T
}

type Manager struct {
	Employee
	Salary         float64
	CommissionRate float64
}

func pay(ps T) (string, float64) {
	var (name string; salary float64)
	switch v := ps.(type) {
	case *Developer:
		name = v.FirstName + v.LastName
		salary = v.HourlyRate * v.HoursWorkedInYear
	case *Manager:
		name = v.FirstName + v.LastName
		salary = v.Salary * (1 + v.CommissionRate)
	default:
		name = "unknown"
		salary = 0
	}
	return name, salary
}

func review(ps *Developer) map[string]T {
	return ps.Review
}

func main() {
	dev := new(Developer)
	dev.FirstName = "H."
	dev.LastName = "Joe"
	dev.HourlyRate = 20
	dev.HoursWorkedInYear = 2000
	dev.Review = map[string]T{"Jane" : "Excellent", "Peter" : 10, "Anna" : "Top"}
	mng := new(Manager)
	mng.FirstName = "M."
	mng.LastName = "Anna"
	mng.Salary = 50000
	mng.CommissionRate = 0.1
	fmt.Println(pay(dev))
	fmt.Println(pay(mng))
	fmt.Println(review(dev))
}
