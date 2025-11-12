package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	UseSyncMap()
	ConcurrencyWaitingForResult()
	ConcurrencyFanOut()
	ConcurrencyFanOut1() // without buffer on channel
	ConcurrencyFanOut2() // buffer on channel
	ConcurrencyFanOut3() // with limit
}

func UseSyncMap() {
	var m sync.Map
	var num int
	var day string

	m.Store("num", 15)
	item, ok := m.Load("num")
	if ok {
		num = item.(int)
		fmt.Println(num)
	}

	m.Store(7, "Friday")
	if item, ok := m.Load(7); ok {
		day = item.(string)
		fmt.Println("The day is: ", day)
	}
	m.Delete("num")
}

func ConcurrencyWaitingForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "sample"
		fmt.Println("Goroutine Finished!")
	}()

	fmt.Println("Waiting...")
	text := <-ch
	fmt.Println(text)
}

var n = 10
var RegisteredUsers = make(chan int, n)

func RegisterUser(userId int) {
	time.Sleep(time.Second)
	RegisteredUsers <- userId
}

func ConcurrencyFanOut() {
	t := time.Now()
	for i := 0; i < n; i++ {
		go RegisterUser(i)
	}

	for n > 0 {
		i := <-RegisteredUsers
		fmt.Println("user ", i, " registered!")
		n--
	}
	fmt.Println(0, time.Since(t))
}

var n1 = 1000000
var RegisteredUsers1 = make(chan int)

func RegisterUser1(userId int) {
	RegisteredUsers1 <- userId
}

func ConcurrencyFanOut1() {
	t := time.Now()
	for i := 0; i < n1; i++ {
		go RegisterUser1(i)
	}

	for n1 > 0 {
		<-RegisteredUsers1
		n1--
	}
	fmt.Println(1, time.Since(t))
}

var n2 = 1000000
var RegisteredUsers2 = make(chan int, n)

func RegisterUser2(userId int) {
	RegisteredUsers2 <- userId
}

func ConcurrencyFanOut2() {
	t := time.Now()
	for i := 0; i < n2; i++ {
		go RegisterUser2(i)
	}

	for n2 > 0 {
		<-RegisteredUsers2
		n2--
	}
	fmt.Println(2, time.Since(t))
}

var n3 = 1000000
var RegisteredUsers3 = make(chan int, n)
var limit = make(chan bool, 3)

func RegisterUser3(userId int) {
	limit <- true
	RegisteredUsers3 <- userId
	<-limit
}

func ConcurrencyFanOut3() {
	t := time.Now()
	for i := 0; i < n3; i++ {
		go RegisterUser3(i)
	}

	for n3 > 0 {
		<-RegisteredUsers3
		n3--
	}
	fmt.Println(3, time.Since(t))
}
