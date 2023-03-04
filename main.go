package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)

	// go leak(&wg)

	// wg.Wait()

	// var wg sync.WaitGroup
	// wg.Add(10)

	// for i := 1; i <= 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// fmt.Println("Done")
	ch1 := make(chan int)
	go sendValue(ch1)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Select timeout ")
	}
}

func sendValue(ch1 chan int) {
	time.Sleep(3 * time.Second)
	ch1 <- 10
}

func leak(wg *sync.WaitGroup) {
	ch := make(chan int)
	ch <- 10
	ch <- 11
	ch <- 12
	go func() {
		val := <-ch
		fmt.Println("Received :", val)
		wg.Done()
	}()
	fmt.Println("Exiting leak method")
	wg.Done()
}
