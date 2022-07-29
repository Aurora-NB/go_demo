package main

import "fmt"

func divide(n int) (res int) {
	res = 100 / n
	return
}

func main() {
	defer fmt.Println("aaaaaaaaaaaaaaaa")
	defer fmt.Println("bbbbbbbbbbbbbbbb")
	defer divide(0)
	defer fmt.Println("cccccccccccccccc")
}
