package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	} else {
		res = a / b
	}
	return
}

func main() {
	// 两种创建 error 的方法
	err1 := fmt.Errorf("%s", "this is normal err1")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is normal err2")
	fmt.Println("err2 = ", err2)

	// error 的使用
	res, err := divide(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
