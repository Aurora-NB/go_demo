package main

import "fmt"

type str string

type Hi struct {
	int
	str
}

func main() {
	a := Hi{1, "hi"}
	fmt.Printf("%+v", a)
}
