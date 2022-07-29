package main

import "fmt"

func fn03(args ...int) {
	fmt.Println("len(args) = ", len(args))
	for i, d := range args {
		fmt.Printf("args[%d] = %d", i, d)
	}
}

func fn03_1(str string, args ...int) {
}

func main() {
	fn03(1, 2, 3, 4, 5)
}
