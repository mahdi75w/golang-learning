package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s, %d years old\n", p.Name, p.Age)
}

type Employee struct {
	Person   // 👈 Embedded field
	Company  string
	Position string
}

func sampleEmbedding() {
	e := Employee{
		Person:   Person{Name: "Hanzo", Age: 30},
		Company:  "Go Corp",
		Position: "Ninja Developer",
	}

	fmt.Println(e.Name)
	fmt.Println(e.Age)

	e.Introduce()

}
