package main

import (
	"fmt"
	"sync"
)

func main() {
	//SellPS5ByDataRace()
	SellPS5ByDataRaceWithMutex()
	//SellPS5ByDataRaceWithBetterMutex()
	SellPS5ByDataRaceWithRMutex()
}

func SellPS5ByDataRace() {
	// if run with -race mode, you got error from this section
	n := 4
	stock := 2
	buyers := 0
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("customer ", i, " entered.")

			if stock <= 0 {
				fmt.Println("not exist")
				return
			}
			fmt.Println("PS5 sold to ", i, " Customer")
			stock--
			buyers++
			println(stock, buyers)
		}(i)

	}
	wg.Wait()
	fmt.Println("Closed")
}

func SellPS5ByDataRaceWithMutex() {
	n := 4
	stock := 2
	buyers := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			fmt.Println("customer ", i, " entered.")

			if stock <= 0 {
				fmt.Println("not exist")
				return
			}
			fmt.Println("PS5 sold to ", i, " Customer")
			stock--
			buyers++
			println(stock, buyers)
		}(i)

	}
	wg.Wait()
	fmt.Println("Closed")
}

func SellPS5ByDataRaceWithBetterMutex() {
	n := 4
	stock := 2
	buyers := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int) {

			defer wg.Done()
			fmt.Println("customer ", i, " entered.")

			mu.Lock()
			outOfStock := stock <= 0
			mu.Unlock()

			if outOfStock {
				fmt.Println("not exist")
				return
			}
			fmt.Println("PS5 sold to ", i, " Customer")
			mu.Lock()
			stock--
			buyers++
			mu.Unlock()
			println(stock, buyers) // ظاهرا این فانکشن همه چیش اوکیه ولی این تیکه که خارج از قفله باعث میشه ریس دیتا به وجود بیاد
			//چون مقدار استوک  و بایر ممکنه توسط گروتین دیگه تغییر کرده باشه
			// برای درک بهتر این پرینت رو کامنت کن و دیگه ارور نداری
		}(i)

	}
	wg.Wait()
	fmt.Println("Closed")
}

func SellPS5ByDataRaceWithRMutex() {
	n := 4
	stock := 2
	buyers := 0
	var wg sync.WaitGroup
	var mu sync.RWMutex

	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int) {

			defer wg.Done()
			fmt.Println("customer ", i, " entered.")

			mu.RLock()
			outOfStock := stock <= 0
			mu.RUnlock()

			if outOfStock {
				fmt.Println("not exist")
				return
			}
			fmt.Println("PS5 sold to ", i, " Customer")
			mu.Lock()
			stock--
			buyers++
			mu.Unlock()
			println(stock, buyers) // ظاهرا این فانکشن همه چیش اوکیه ولی این تیکه که خارج از قفله باعث میشه ریس دیتا به وجود بیاد
			//چون مقدار استوک  و بایر ممکنه توسط گروتین دیگه تغییر کرده باشه
			// برای درک بهتر این پرینت رو کامنت کن و دیگه ارور نداری
		}(i)

	}
	wg.Wait()
	fmt.Println("Closed")
}
