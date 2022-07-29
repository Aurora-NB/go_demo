package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	s := make(chan int)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(5 * time.Second):
				fmt.Println("超时")
				s <- 1
				runtime.Goexit()
			}
		}
	}()
	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<-s
	fmt.Println("程序结束")
}
