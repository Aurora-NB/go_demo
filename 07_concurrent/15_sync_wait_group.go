package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHello() {
	fmt.Println("hello")
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sayHello()
	}
	wg.Wait()
}
