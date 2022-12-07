package shape

import "fmt"

type Shape interface {
	area() float64
	name() string
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Square struct {
	Side float64
}

func PrintShapeDetails(shape ...Shape) {
	for _, item := range shape {
		fmt.Printf("The area of %s is: %.2f\n", item.name(), item.area())
	}
}

func (t Triangle) area() float64 {
	return t.Base * t.Height / 2
}
func (t Triangle) name() string {
	return "Triangle"
}

func (rec Rectangle) area() float64 {
	return rec.Length * rec.Width
}
func (rec Rectangle) name() string {
	return "Rectangle"
}

func (sq Square) area() float64 {
	return sq.Side * sq.Side
}
func (sq Square) name() string {
	return "Square"
}
