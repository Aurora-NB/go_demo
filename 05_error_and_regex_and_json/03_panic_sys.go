package main

import "fmt"

func a2() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaa")
}
func b2(d int) {
	a := 1 / d
	fmt.Println(a)
}
func c2() {
	fmt.Println("ccccccccccccccccccccc")
}

func main() {
	a2()
	b2(0)
	c2()
}
