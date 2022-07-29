package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			fmt.Println(m.Load(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
