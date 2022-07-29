package main

import (
	"fmt"
	"sync"
	"time"
)

var wg3 sync.WaitGroup
var rwLock sync.RWMutex
var x = 0

func read() {
	defer wg3.Done()
	rwLock.RLock()
	fmt.Println(x)
	rwLock.RUnlock()
}
func write() {
	defer wg3.Done()
	rwLock.RLock()
	x++
	rwLock.RUnlock()

}
func main() {
	begin := time.Now()
	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go write()
	}
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go read()
	}
	wg3.Wait()
	end := time.Now()
	fmt.Println(end.Sub(begin))
}
