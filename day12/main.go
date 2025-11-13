package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	WaitingForChannel()
	pooling()
	drop()
	timout()

	ErrorSampleMain()
	OopMainFunc()
}

type User struct {
	Name string
	Age  int
}

func WaitingForChannel() {
	ch := make(chan User)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		user := <-ch
		time.Sleep(2 * time.Second)
		fmt.Println(user.Name, "is Registered!")
		wg.Done()
	}()

	ch <- User{
		Name: "Mahdi",
		Age:  12,
	}
	fmt.Println("data send")
	wg.Wait()
	fmt.Println("----------------------------------")
}

func pooling() {
	n := runtime.NumCPU()
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			for name := range ch { // its work until channel become close
				fmt.Println("goroutine", i, "printing name:", name)
			}
			wg.Done()
		}(i)
	}
	names := []string{
		"Alice",
		"mahdi",
		"mary",
		"michael",
		"Ali",
	}

	for _, name := range names {
		ch <- name
	}
	close(ch)
	wg.Wait()
	fmt.Println("all goroutines done", n)
	fmt.Println("----------------------------------")
}

func drop() {
	n := runtime.NumCPU()
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			for request := range ch { // its work until channel become close
				fmt.Println("worker", i, "revived request of ", request)
			}
			wg.Done()
		}(i)
	}

	for j := 0; j <= 20; j++ {
		select {
		case ch <- j:
		default:
			fmt.Println("drop channel full", j)
		}
	}
	close(ch)
	wg.Wait()
	fmt.Println("all goroutines done", n)
	fmt.Println("----------------------------------")
}

func timout() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 1) // if you increase up to 2 second finished with timeout
		ch <- "its worked"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timed out")
	}
	fmt.Println("----------------------------------")
}
