package main

import "fmt"

func main() {
	var p *int
	a := 1
	p = &a
	*p = 666
	fmt.Println(*p)

	p1 := new(int)
	*p1 = 777
	fmt.Println(*p1)
}
