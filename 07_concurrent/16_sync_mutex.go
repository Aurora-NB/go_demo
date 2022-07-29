package main

import (
	"fmt"
	"sync"
)

var x1 = 0
var wg1 sync.WaitGroup

var x2 = 0
var wg2 sync.WaitGroup
var lock sync.Mutex

func add1() {
	defer wg1.Done()
	for i := 0; i < 5000; i++ {
		x1 = x1 + 1
	}

}

func add2() {
	defer wg1.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x1 = x1 + 1
		lock.Unlock()
	}

}

func main() {
	// 不通过锁结果可能不为 10000
	wg1.Add(2)
	go add1()
	go add1()
	wg1.Wait()
	fmt.Println("x1:", x1)

	// 加锁结果一定为 10000
	wg2.Add(2)
	go add2()
	go add2()
	wg2.Wait()
	fmt.Println("x2", x2)
}
