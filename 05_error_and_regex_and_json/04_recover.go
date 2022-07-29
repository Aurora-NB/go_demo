package main

import "fmt"

func a3() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaa")
}
func b3(d int) {
	// recover 一定要在 defer 中使用
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	a := 1 / d
	fmt.Println(a)
}
func c3() {
	fmt.Println("ccccccccccccccccccccc")
}

func main() {
	a3()
	b3(0)
	c3()
}
