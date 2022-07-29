package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Key string

var wg2 sync.WaitGroup

func worker(ctx context.Context) {
	s, ok := ctx.Value(Key("Work")).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}

LOOP:
	for {
		fmt.Println("code:", s)
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("over")
	wg2.Done()
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, Key("Work"), "122334")
	wg2.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg2.Wait()

}
