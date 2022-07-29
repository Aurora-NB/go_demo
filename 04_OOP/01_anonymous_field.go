package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender byte
}

type Student struct {
	Person // 匿名字段，相当于继承
	addr   string
	age    int
}

func main() {
	s := Student{
		addr: "浙江",
		age:  15,
		Person: Person{
			age: 20,
		},
	}
	fmt.Printf("%+v", s)
}
