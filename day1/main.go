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
	find_min_and_max_with_builtin_func()
	casting_int_and_float()
	sample_for_iota()
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

func find_min_and_max_with_builtin_func() {
	a, b := 10, 20

	// third parameter in min and max are optional and return static number that you must return if other  is less
	fmt.Println(max(a, b))
	fmt.Println(min(a, b))
}

func casting_int_and_float() {
	a, b := 12, 3.14
	fmt.Println(a, b)
	fmt.Println(float32(a), int8(b))
}

func sample_for_iota() {
	const (
		status_1 = iota + 1
		status_2
		status_3
		status_4
	)

	fmt.Println("status_3: ", status_3)

	const (
		reason_1 = iota + 1
		_
		_
		reason_4
	)
	fmt.Println("reason_4: ", reason_4)
}
