package main

import "fmt"

func a() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaa")
}
func b() {
	panic("出现异常")
}
func c() {
	fmt.Println("ccccccccccccccccccccc")
}

func main() {
	a()
	b()
	c()
}
