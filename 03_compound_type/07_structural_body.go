package main

import "fmt"

type Student struct {
	id       int
	name     string
	gender   byte
	age      int
	handsome bool
}

func main() {
	// 必须赋值所有属性
	var s1 Student
	s1 = Student{1, "wang", 'm', 20, true}
	fmt.Println(s1)

	// 未被赋值的属性设为0
	s2 := Student{age: 20}
	fmt.Println(s2)

	var s3 Student
	s3.id = 1
	s3.gender = 'f'
	s3.age = 19
	s3.name = "zhang"
	fmt.Println(s3)

	s4 := new(Student)
	fmt.Println(s4)

	// 赋值操作是值传递
	var tmp Student
	tmp = s2
	fmt.Printf("%p\n%p", &tmp, &s2)
}
