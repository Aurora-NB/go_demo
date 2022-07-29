package main

import "fmt"

// 空接口可以接收任意类型的参数
func myPrint(args ...interface{}) {
	for _, v := range args {
		fmt.Printf("%v ", v)
	}
	println()
}

func main() {
	myPrint(1, true, "hi", 'k')
}
