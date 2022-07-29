package main

import "fmt"

type S_4_7 struct {
}

func (s S_4_7) fn(i int) {
	fmt.Println("hi", i)
}

func main() {

	s := S_4_7{}
	// 方法值
	fn1 := s.fn
	fn1(12)
	// 方法表达式
	fn2 := S_4_7.fn
	fn2(s, 123)
}
