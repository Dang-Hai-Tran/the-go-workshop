package main

import (
	"fmt"
)
type Action interface {
	Stop() string
	TurnLeft() string
}
type Car struct {
	Color string
	Material string
}

// Implement method in golang
// Car have method Stop() so Car is an Action interface
func (c Car) Stop() string {
	return "Stop!"	
}
func (c Car) TurnLeft() string {
	return "TurnLeft!"
}

// Implement inheritance in golang
type Renaut struct {
	Car
	Branch string
}

func turnLeft(oj Action) {
	fmt.Println(oj.TurnLeft())
}

func main() {
	newCar := new(Car)
	newCar.Color = "red"
	newCar.Material = "plastic"
	fmt.Println(newCar.Stop())
	fmt.Println(newCar.Color, newCar.Material)
	turnLeft(newCar)
	newRenault := Renaut{Car{"black", "steel"}, "Renaut"}
	fmt.Println(newRenault.Color)
	fmt.Println(newRenault.Stop())

}
