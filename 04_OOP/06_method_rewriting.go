package main

import "fmt"

type Person_4_6 struct {
	id   int
	name string
}

func (p Person_4_6) fn() {
	fmt.Printf("id=%d,name=%s", p.id, p.name)
}

func (p Man_4_6) fn() {
	fmt.Printf("Man = %+v", p)
}

type Man_4_6 struct {
	Person_4_6
	age int
}

func main() {
	man := Man_4_6{
		Person_4_6: Person_4_6{
			id:   0,
			name: "xiao ming",
		},
		age: 0,
	}
	man.fn()
}
