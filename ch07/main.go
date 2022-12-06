package main

import (
	"fmt"
)

type Car struct {
	Color string
	Material string
}

// Implement method in golang
func (c Car) Stop() string {
	return "Stop!"	
}

// Implement inheritance in golang
type Renaut struct {
	Car
	Branch string
}

func main() {
	newCar := Car{"red", "steel"}
	fmt.Println(newCar.Stop())
	fmt.Println(newCar.Color, newCar.Material)
	newRenault := Renaut{Car{"black", "steel"}, "Renaut"}
	fmt.Println(newRenault.Color)
	fmt.Println(newRenault.Stop())

}
