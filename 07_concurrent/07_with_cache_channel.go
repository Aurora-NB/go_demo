package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("i = ", i)
			ch <- i
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("num = ", num)
	}
}
