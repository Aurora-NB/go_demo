package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU()
	n := runtime.GOMAXPROCS(1)
	fmt.Println(num, n)

	//go func() {
	//	for {
	//		fmt.Print(0)
	//	}
	//}()
	//
	//for {
	//	fmt.Print(1)
	//}

}
