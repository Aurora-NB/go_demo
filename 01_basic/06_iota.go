package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1)
	const (
		a2         = iota
		b2, c2, d2 = iota, iota, iota
		e2         = iota
	)
	fmt.Println(a2, b2, c2, d2, e2)

}
