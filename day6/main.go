package main

import (
	_ "encoding"       // import package without usage
	"example/package1" // why start by example? because i init project with example name
	myname "fmt"
)

var numberOfDay int

func init() {
	numberOfDay = 6
}

func main() {
	myname.Println("Day ", numberOfDay)
	package1.SayHello()

	defaine_Array()
	calculate_len_of_array()
	print_all_of_index_in_array()
	define_slice()
	work_with_append()
	check_increase_len_And_capacity_slice()
	slice_on_array()
	copy_of_slice()
	another_example_of_copy()
	range_in_array()
	clear_slice()
}

func defaine_Array() {
	var my_array [5]int // default value is 0
	myname.Println(my_array)

	my_Str_array := [5]string{} // default value is ""
	myname.Println(my_Str_array)
}

func calculate_len_of_array() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myname.Println(len(numbers))
}

func print_all_of_index_in_array() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index := 0; index < len(numbers); index++ {
		myname.Println(numbers[index])
	}
}

func define_slice() {
	number := make([]int, 8, 16)
	myname.Println("len ", len(number))
	myname.Println("cap ", cap(number))
}

func work_with_append() {
	number := make([]int, 8, 16)
	new_numner := append(number, 1, 2, 5, 8)
	myname.Println(number)
	myname.Println(new_numner)
}

func check_increase_len_And_capacity_slice() {
	numbers := make([]int, 3, 6)
	for i := 0; i < 8; i++ {
		myname.Println(len(numbers), cap(numbers), numbers)
		numbers = append(numbers, i)
	}
}

func slice_on_array() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myname.Println(numbers[2:5])
}

func copy_of_slice() {
	nums := make([]int, 4)
	myname.Println(nums)
	copy(nums, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	myname.Println(nums)
}

func another_example_of_copy() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := copy(numbers, []int{11, 12})
	myname.Println(n, numbers)
}

func range_in_array() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range numbers {
		myname.Println(index, " : ", value)
	}
}

func clear_slice() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myname.Println(numbers)
	clear(numbers)
	myname.Println(numbers)
}
