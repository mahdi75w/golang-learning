package main

import (
	"fmt"
	"math"
)

func OopMainFunc() {
	r := Rectangle{
		width:  10,
		height: 11,
	}
	c := Circle{
		radius: 10,
	}

	PrintArea(r)
	PrintArea(c)
}

type OopUser struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *OopUser {
	return &OopUser{
		Name: name,
		Age:  age,
	}
}

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func PrintArea(shape Shape) {
	fmt.Println(shape.Area())
}
