package main

import "fmt"

func fn04() int {
	return 1
}

func fn04_1() (res int) {
	return 1
}

func fn04_2() (res int) {
	res = 1
	return
}

func main() {
	a := fn04()
	fmt.Println(a)

	b := fn04_1()
	fmt.Println(b)

	c := fn04_2()
	fmt.Println(c)
}
