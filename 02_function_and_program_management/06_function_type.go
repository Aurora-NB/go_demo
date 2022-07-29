package main

import "fmt"

type FunctionType func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func main() {
	var fn FunctionType

	fn = Add
	res1 := fn(10, 5)
	fmt.Println(res1)

	fn = Minus
	res2 := fn(10, 5)
	fmt.Println(res2)

	fn = func(i int, i2 int) int {
		return i / i2
	}
	res3 := fn(10, 5)
	fmt.Println(res3)

}
