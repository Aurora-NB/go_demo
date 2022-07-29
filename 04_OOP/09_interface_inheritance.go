package main

import "fmt"

type Hunmaner_4_9 interface {
	walk()
}
type Studenter_4_9 interface {
	Hunmaner_4_9 // 通过匿名字段进行继承
	sing()
}
type Student_4_9 struct {
	name string
	age  int
}

func (s Student_4_9) walk() {
	fmt.Printf("%+v 正在走\n", s)
}
func (s Student_4_9) sing() {
	fmt.Printf("%+v 正在唱歌\n", s)
}

func main() {
	var i Studenter_4_9
	s := Student_4_9{"xiaoming", 18}
	i = s
	i.walk()
	i.sing()
}
