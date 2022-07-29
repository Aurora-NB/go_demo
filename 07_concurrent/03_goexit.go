package main

import (
	"fmt"
	"runtime"
)

func fn() {
	defer fmt.Println("ccc")
	runtime.Goexit()
	fmt.Println("ddd")
}

func main() {

	go func() {
		fmt.Println("aaa")
		fn()
		fmt.Println("bbb")
	}()
	for {
	}
}
