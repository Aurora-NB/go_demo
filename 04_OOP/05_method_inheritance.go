package main

import "fmt"

type Person_4_5 struct {
	id   int
	name string
}

func (p Person_4_5) fn() {
	fmt.Printf("id=%d,name=%s", p.id, p.name)
}

type Man_4_5 struct {
	Person_4_5
	age int
}

func main() {
	man := Man_4_5{
		Person_4_5: Person_4_5{
			id:   0,
			name: "xiao ming",
		},
		age: 0,
	}
	man.fn()
}
