package main

import (
	"fmt"
	"time"
)

func main() {
	go printLetters() // for check speed remove and add go in run code
	printNumbers()
	fmt.Println("Finished")
}

func printLetters() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Print(string(i), " ")
		time.Sleep(time.Millisecond * 50)
	}
}

func printNumbers() {
	for i := 1; i <= 50; i++ {
		fmt.Print(i, " ")
		time.Sleep(time.Millisecond * 50)
	}
}
