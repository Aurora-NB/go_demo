package main

import (
	"fmt"
	"sync"
)

// 类似于单例模式下的 DCL
var loadResourceOnce sync.Once
var chani = make(chan int)

func loadResource(x int) {
	// 这里是一些只需要被执行一次用于加载资源的代码
	fmt.Println("资源已加载", x)
	chani <- 1
}

func main() {
	x := 1
	var closer = func() func() {
		return func() {
			loadResource(x)
		}
	}()
	// 模拟 100 个协程同时访问
	for i := 0; i < 100; i++ {
		go func() {
			loadResourceOnce.Do(closer)
		}()
	}

	<-chani
}
