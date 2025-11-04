package main

import (
	"fmt"
)

func main() {
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(1, 2))
	//-----------
	use_inner_func_by_anonymous_function()
	// ----------
	func() {
		fmt.Println("without variable ")
	}()
	// ----------
	use_variable_in_up_scope()
	// ----------
	use_new_counter()
	// ----------
	myPrint(hello("mahdi"))
	// ----------
	fmt.Println(applyFunc(5, double))
	fmt.Println(applyFunc(5, square))
	// ----------
	use_defer_in_simple_way()

}

func use_inner_func_by_anonymous_function() {
	add := func() {
		fmt.Println("i am in inner func")
	}

	fmt.Println("i am in main func")
	add()
}

func use_variable_in_up_scope() {
	i := 0
	next := func() int {
		i++
		return i
	}
	fmt.Println(next())
	fmt.Println(next())
}

func new_counter() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func use_new_counter() {
	player_1 := new_counter()
	player_2 := new_counter()
	player_3 := new_counter()

	fmt.Println("player1 :", player_1())
	fmt.Println("player2 :", player_2())
	fmt.Println("player1 :", player_1())
	fmt.Println("player3 :", player_3())
}

func hello(name string) string {
	return "Hello, " + name
}

func myPrint(s string) {
	fmt.Println(s)
}

func double(num int) int { return num * 2 }

func square(num int) int { return num * num }

func applyFunc(n int, f func(int) int) int {
	result := f(n)
	return result
}

func use_defer_in_simple_way() {
	defer fmt.Println("I am in simple way")
	defer fmt.Println("I am in simple way2")
	defer fmt.Println("I am in simple way3")

	fmt.Println("I am in simple way4")
	fmt.Println("I am in simple way5")
}
