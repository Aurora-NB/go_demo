package main

import "fmt"

func fn05() (int, int, int) {
	return 1, 2, 3
}

func fn05_1() (a, b, c int) {
	return 1, 2, 3
}

func fn05_2() (a, b, c int) {
	a = 1
	b = 2
	c = 3
	return
}

func main() {
	a1, a2, a3 := fn05()
	fmt.Println(a1, a2, a3)

	b1, b2, b3 := fn05_1()
	fmt.Println(b1, b2, b3)

	c1, c2, c3 := fn05_2()
	fmt.Println(c1, c2, c3)
}
