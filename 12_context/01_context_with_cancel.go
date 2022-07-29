package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doSomething(ctx context.Context) {
	defer wg.Done()
ForLoop:
	for {
		fmt.Println("hello, world")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break ForLoop
		default:

		}
	}
}

func main() {
	wg.Add(1)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go doSomething(ctx)
	time.Sleep(time.Second * 3)
	cancelFunc()
	wg.Wait()
}
