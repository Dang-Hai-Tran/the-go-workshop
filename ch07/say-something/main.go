package main

import (
	"fmt"
)

type Cat struct {
}
type Dog struct {
}
type Person struct {
	Name string
}

type Speak interface {
	Speak() string
}

func (c Cat) Speak() string {
	return "Meow Meow"
}
func (c Dog) Speak() string {
	return "Woof Woof"
}
func (c Person) Speak() string {
	return fmt.Sprintf("Hi, my name is %v", c.Name)
}
func saySomeThing(c ...Speak) {
	for _, v := range c {
		fmt.Println(v.Speak())
	}
}

func main() {
	cat := new(Cat)
	dog := new(Dog)
	nam := new(Person)
	nam.Name = "Nam"
	saySomeThing(cat, dog, nam)
}
