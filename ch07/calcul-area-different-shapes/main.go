package main

import (
	"fmt"
)

type Rectangle struct {
	name   string
	height float64
	width  float64
}

type Triangle struct {
	name   string
	base   float64
	height float64
}

func (shape Rectangle) Area() float64 {
	return shape.height * shape.width
}

func (shape Triangle) Area() float64 {
	return shape.height * shape.base / 2
}

func (shape Rectangle) Name() string {
	return shape.name
}
func (shape Triangle) Name() string {
	return shape.name
}

type Shape interface {
	Area() float64
	Name() string
}

func printArea(shape ...Shape) {
	for _, v := range shape {
		fmt.Printf("This %v area is %v\n", v.Name(), v.Area())
	}
}

func main() {
	rec := new(Rectangle)
	rec.name = "rectangle"
	rec.height = 5
	rec.width = 10.4
	tri := new(Triangle)
	tri.name = "triangle"
	tri.base = 3
	tri.height = 4
	printArea(rec, tri)
}
