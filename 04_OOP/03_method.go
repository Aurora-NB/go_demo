package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) printInfo() {
	fmt.Println(h)
}

func main() {
	h := Human{"king", 22}
	h.printInfo()
}
