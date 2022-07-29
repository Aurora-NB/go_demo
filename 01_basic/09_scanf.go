package main

import "fmt"

func main() {
	var a int
	fmt.Print("请输入 a:")
	fmt.Scanf("%d", &a)
	fmt.Printf("%d\n", a)

	var b int
	fmt.Print("请输入 b:")
	fmt.Scan(&b)
	fmt.Print(b)
}
