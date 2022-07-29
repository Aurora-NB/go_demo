package main

import "fmt"

type Tea struct {
}

func (t Tea) byVar() {
	fmt.Println("byVar")
}

func (t *Tea) byPointer() {
	fmt.Println("byPointer")
}

func main() {
	v := Tea{}
	p := &Tea{}
	// 指针和普通变量都会自动转换
	v.byVar()
	v.byPointer()
	p.byVar()
	p.byPointer()
}
