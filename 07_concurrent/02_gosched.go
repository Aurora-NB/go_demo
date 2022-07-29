package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 5; i++ {
		// 让出时间片，让其他协程先运行
		runtime.Gosched()
		fmt.Println("hello")
	}
}
