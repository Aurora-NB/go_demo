package main

import "fmt"

func fn() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	fn := fn()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}
