package main

import "fmt"

func main() {
	a := 1
	b := 'a'
	c := true
	d := "hello"
	e := 3.1415926
	fmt.Printf("%T,%T,%T,%T,%T\n", a, b, c, d, e)
	fmt.Printf("%d,%c,%t,%s,%f\n", a, b, c, d, e)
	fmt.Printf("%v,%v,%v,%v,%v\n", a, b, c, d, e)

}
