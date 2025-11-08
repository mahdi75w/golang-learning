package main

import (
	"fmt"
	"time"
)

func main() {
	go printLetters()
	go printNumbers()
	fmt.Println("Finished")
	// you can see printed finished without print letter or number -> because not blocked
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
