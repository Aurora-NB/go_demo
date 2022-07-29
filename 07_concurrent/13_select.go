package main

import "fmt"

func consumer_7_13(ch <-chan int, flag chan<- int) {
	fmt.Println("consumer")
	for {
		for i := 0; i < 10; i++ {
			fmt.Print(<-ch, " ")
		}
		fmt.Println()
		flag <- 1
		flag <- 1
		return
	}
}

func producer_7_13(ch chan<- int, flag <-chan int) {
	fmt.Println("producer")
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-flag:
			return
		}
	}
}

func main() {
	ch := make(chan int)
	flag := make(chan int)
	go consumer_7_13(ch, flag)
	go producer_7_13(ch, flag)
	<-flag
}
