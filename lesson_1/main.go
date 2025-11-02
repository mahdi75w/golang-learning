package main

import (
	"fmt"
)

func main() {
	normaly_hello_world()
	multi_variable_in_block()
	multi_variable_in_block_without_point_to_type()
	create_variable_without_point_to_var()
	new_varibale_define()
}

func normaly_hello_world() {
	fmt.Println("hello world!")
}

func multi_variable_in_block() {
	var (
		age  int    = 29
		name string = "mahdi"
	)
	fmt.Println("my name is ", name, " and my age is", age)
}

func multi_variable_in_block_without_point_to_type() {
	var (
		age  = 29
		name = "mahdi"
	)

	fmt.Println("my name is ", name, " and my age is", age)
}

func create_variable_without_point_to_var() {
	// only available in func
	name := "mahdi"
	age := 29
	fmt.Println("my name is ", name, " and my age is", age)
}

func new_varibale_define() {
	first_name, age := "mahdi", 29
	fmt.Println("my name is ", first_name, " and my age is", age)
}
