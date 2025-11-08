package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		printLetters()
		wg.Done()
	}()

	go func() {
		printNumbers()
		wg.Done()
	}()
	wg.Wait()
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
