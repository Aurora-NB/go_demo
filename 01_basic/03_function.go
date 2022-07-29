package main

import "fmt"

func fn() (a, b, c int) {
	return 1, 2, 3
}
func main() {
	var c int = 0
	_, b, c := fn()
	fmt.Println(b, c)

}
