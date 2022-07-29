package main

import "fmt"

func Calculate(a, b int, fn func(int, int) int) (res int) {
	res = fn(a, b)
	return
}
func main() {
	a := Calculate(1, 2, func(i int, i2 int) int {
		return i + i2
	})
	fmt.Println(a)

	b := Calculate(1, 2, func(i int, i2 int) int {
		return i - i2
	})
	fmt.Println(b)

	c := Calculate(1, 2, func(i int, i2 int) int {
		return i * i2
	})
	fmt.Println(c)

	d := Calculate(1, 2, func(i int, i2 int) int {
		return i / i2
	})
	fmt.Println(d)

}
