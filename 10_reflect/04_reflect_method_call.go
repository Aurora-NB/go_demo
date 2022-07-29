package main

import (
	"fmt"
	"reflect"
)

type User_10_4 struct {
	Name string
	Age  int
	Id   int
}

func (u User_10_4) Hello(s string) {
	fmt.Println(s, "hello")
}

func main() {
	o := User_10_4{
		Name: "king",
		Age:  13,
		Id:   3,
	}

	re := reflect.ValueOf(o)
	helloFunc := re.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("jin")}
	if helloFunc.IsValid() {
		helloFunc.Call(args)
	}
}
