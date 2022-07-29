package main

import "fmt"

type MyType interface {
	~int | ~int64 | ~string
}

func Print[T MyType](in ...T) {
	fmt.Println(in)
}

func main() {
	Print("abc")
	Print(1, 2)
}