package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := [5]int{6, 7, 8, 9, 0}
	fmt.Println(b)

	c := [5]int{6, 7, 8}
	fmt.Println(c)

	d := [5]int{2: 10, 4: 15}
	fmt.Println(d)
}
