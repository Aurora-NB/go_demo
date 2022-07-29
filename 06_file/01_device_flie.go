package main

import (
	"fmt"
)

func main() {
	// 标准输出关闭后无法再输出
	//err := os.Stdout.Close()
	//if err != nil {
	//	return
	//}
	fmt.Println("out")

	// 标准输入关闭后无法在输出
	//err2 := os.Stdin.Close()
	//if err2 != nil {
	//	return
	//}
	var i int
	_, _ = fmt.Scanf("%d", &i)
}
