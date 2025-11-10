package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	WithoutSelect()
	UseSelect()
	CloseCoffeeShopIn3SecondAfterLastChannel()
	CloseCoffeeShopIn3SecondAfterLastChannelWithTimeAfter()
	CloseCoffeeShopIn3SecondAndWorkWithTicker()
}

func WithoutSelect() {
	coffeeCh := make(chan string)
	teaCh := make(chan string)
	foodCh := make(chan string)

	go func() {
		time.Sleep(time.Second)
		teaCh <- "Green Tea"
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		coffee := <-coffeeCh
		fmt.Println("Recived a coffee order: ", coffee)
		wg.Done()
	}()

	go func() {
		tea := <-teaCh
		fmt.Println("Recived a tea order: ", tea)
		wg.Done()
	}()

	go func() {
		food := <-foodCh
		fmt.Println("Recived a food order: ", food)
		wg.Done()
	}()

	wg.Wait()
}

func UseSelect() {
	coffeeCh := make(chan string)
	teaCh := make(chan string)
	foodCh := make(chan string)

	go func() {
		time.Sleep(time.Second)
		teaCh <- "Green Tea"
	}()

	select {
	case coffee := <-coffeeCh:
		fmt.Println("Recived a coffee order: ", coffee)
	case tea := <-teaCh:
		fmt.Println("Recived a tea order: ", tea)
	case food := <-foodCh:
		fmt.Println("Recived a food order: ", food)

	}

}

func CloseCoffeeShopIn3SecondAfterLastChannel() {
	coffeeCh := make(chan string)
	teaCh := make(chan string)
	foodCh := make(chan string)
	quit := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		teaCh <- "Green Tea"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		foodCh <- "Burger"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		quit <- true
	}()

	defer fmt.Println("shop is closed!")

	for {
		select {
		case coffee := <-coffeeCh:
			fmt.Println("Recived a coffee order: ", coffee)
		case tea := <-teaCh:
			fmt.Println("Recived a tea order: ", tea)
		case food := <-foodCh:
			fmt.Println("Recived a food order: ", food)
		case <-quit:
			return
		default:
		}
	}
}

func CloseCoffeeShopIn3SecondAfterLastChannelWithTimeAfter() {
	coffeeCh := make(chan string)
	teaCh := make(chan string)
	foodCh := make(chan string)
	quit := time.After(time.Second * 3)

	go func() {
		time.Sleep(time.Second)
		teaCh <- "Green Tea"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		foodCh <- "Burger"
	}()

	defer fmt.Println("shop is closed!")

	for {
		select {
		case coffee := <-coffeeCh:
			fmt.Println("Recived a coffee order: ", coffee)
		case tea := <-teaCh:
			fmt.Println("Recived a tea order: ", tea)
		case food := <-foodCh:
			fmt.Println("Recived a food order: ", food)
		case <-quit:
			return
		default:
		}
	}
}

func CloseCoffeeShopIn3SecondAndWorkWithTicker() {
	coffeeCh := make(chan string)
	teaCh := make(chan string)
	foodCh := make(chan string)
	quit := time.After(time.Second * 3)
	tickerCh := time.Tick(time.Millisecond * 500)

	go func() {
		time.Sleep(time.Second)
		teaCh <- "Green Tea"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		foodCh <- "Burger"
	}()

	defer fmt.Println("shop is closed!")

	for {
		select {
		case coffee := <-coffeeCh:
			fmt.Println("Recived a coffee order: ", coffee)
		case tea := <-teaCh:
			fmt.Println("Recived a tea order: ", tea)
		case food := <-foodCh:
			fmt.Println("Recived a food order: ", food)
		case <-tickerCh:
			fmt.Println("Tick")
		case <-quit:
			return
		default:
		}
	}
}
