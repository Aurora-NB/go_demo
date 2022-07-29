package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func connDatabase(ctx context.Context) {
	defer wg1.Done()
	fmt.Println("连接数据库...")
	time.Sleep(60 * time.Millisecond) // 假设要十毫秒
	select {
	case <-ctx.Done():
		fmt.Println("连接超时...")
	default:
		fmt.Println("连接成功...")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50) // 超过50ml连接失败
	defer cancel()
	wg1.Add(1)
	go connDatabase(ctx)
	wg1.Wait()
}
