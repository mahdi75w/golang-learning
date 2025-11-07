package main

import (
	_ "encoding"       // import package without usage
	"example/package1" // why start by example? because i init project with example name
	mynme "fmt"
)

var numberOfDay int

func init() {
	numberOfDay = 6
}

func main() {
	mynme.Println("Day ", numberOfDay)
	package1.SayHello()
}
