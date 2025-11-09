package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	normalDoTask()
	runWithGoNormalDoTask()
	WorkWithChannel()
	ChannelWithExtraCapacity()
	ManageNumberOfGoroutine()
	iterateChannelType1()
	iterateChannelType2()
}

func normalDoTask() {
	const n = 10
	t := time.Now()
	for i := 0; i < n; i++ {
		DoTask(i, time.Millisecond*time.Duration(i))
	}
	fmt.Println(time.Since(t))
}

func DoTask(i int, t time.Duration) {
	time.Sleep(t)
	fmt.Println("task ", i, " finished!")
}

func runWithGoNormalDoTask() {
	const n = 10
	wg := sync.WaitGroup{}
	wg.Add(n)

	t := time.Now()
	for i := 0; i < n; i++ {
		go func(i int) {
			DoTask(i, time.Millisecond*time.Duration(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(t))
}

func HelloWorld(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello World"
}

func WorkWithChannel() {
	c := make(chan string)
	go HelloWorld(c)
	fmt.Println("waiting...")
	message := <-c
	fmt.Println(message)
}

func ChannelWithExtraCapacity() {
	c := make(chan string, 2)
	c <- "first data"
	c <- "second data"
	fmt.Println(<-c)
}

func HeavyTask(b chan int) {
	time.Sleep(time.Second)
	i := <-b
	fmt.Println("task ", i, " finished!")
}

func ManageNumberOfGoroutine() {
	c := make(chan int, 3)
	for i := 0; i < 100; i++ {
		c <- i
		fmt.Println("Sent", i)
		go HeavyTask(c)
	}
}

func iterateChannelType1() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for n := range c {
		fmt.Println(n)
	}

}

func iterateChannelType2() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	for i := 1; i < 3; i++ {
		fmt.Println(<-c)
	}

}
