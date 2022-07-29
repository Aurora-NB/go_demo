package main

import "fmt"

// make用于生成切片
// copy用于复制切片
// append当切片容量不够时，底层创建一个新的数组，长度为原来的2倍，超出1000个数时长度为原来的1.25倍
// 函数传递的是切片的复制

func testSlice(a []int) []int {
	a[0] = 10
	return a

}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a
	// 数组是值传递
	fmt.Printf("%p\n%p\n", &a, &t)

	s := a[0:3:5]
	testSlice(s)
	fmt.Println(a, s)

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = a[1:4:5]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// 切片是地址传递
	var k = s
	fmt.Printf("%p\n%p", k, s)

}
