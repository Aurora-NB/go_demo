package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 0
	// 修改值必须使用地址
	value := reflect.ValueOf(&x)
	value.Elem().SetInt(666)

	fmt.Println(x)
}
