package main

import "fmt"

func main() {
	// 双向管道可以转换为单向管道，反之则不可以
	ch := make(chan int)
	var writeCh chan<- int = ch
	var readCh <-chan int = ch
	go func() {
		writeCh <- 1
	}()
	i := <-readCh
	fmt.Println(i)
}
