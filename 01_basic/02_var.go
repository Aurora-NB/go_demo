package main

import "fmt"

func main() {
	var a int
	a = 1
	var b int = 2
	var c = 3
	d := 4
	fmt.Println(a, b, c, d)

	var i, j = 1, 2
	i, j = j, i
	fmt.Println(i, j)
}
