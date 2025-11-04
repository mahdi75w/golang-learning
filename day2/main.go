package main

import "fmt"

func main() {
	switch_example()
	advance_switch_example()
	switch_without_default_condition()
	impliment_while_with_for()
	fmt.Println(area_and_perimeter(50, 38))
	fmt.Println(area_and_perimeter_with_return_by_name(12, 19))
}

func switch_example() {
	day := 4

	switch day {
	case 1:
		fmt.Println("day is ", 1)
	case 2:
		fmt.Println("day is ", 2)
	case 3:
		fmt.Println("day is ", 3)
	default:
		fmt.Println("day is not found")
	}
}

func advance_switch_example() {
	// Unlike Python, there is no need to define a list to check multiple items at the same time.
	food := "Tea"
	switch food {
	case "Coca", "Water", "Tea", "Coffee":
		fmt.Println("you are order a drink")
	case "pizza", "burger", "sandwich":
		fmt.Println("you are order a main food")
	case "salad", "french fries", "yogurt":
		fmt.Println("you are order a side food")
	default:
		fmt.Println("unknown food")
	}
}

func switch_without_default_condition() {
	food := "pizza"
	amount := 200000

	switch {
	case amount > 100000:
		fmt.Println("discount is 30% so amount is ", float64(amount)*0.7)
	case food == "pizza":
		fmt.Println("discount is 25% so amount is ", float64(amount)*0.85)
	case food == "burger":
		fmt.Println("discount is 10% so amount is ", float64(amount)*0.90)
	}
}

func impliment_while_with_for() {
	// default value in int is 0
	var i int
	for i < 5 {
		fmt.Println("i is ", i)
		i++
	}
}

func area_and_perimeter(a, b int) (int, int) {
	return 2 * (a + b), a * b
}

func area_and_perimeter_with_return_by_name(a, b int) (area int, perimeter int) {
	area = 2 * (a + b)
	perimeter = a * b
	return
}
