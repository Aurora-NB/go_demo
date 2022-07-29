package main

import (
	"fmt"
	"reflect"
)

type User_10_3 struct {
	Name string
	Age  int
	Id   int
}

func main() {
	o := User_10_3{
		Name: "King",
		Age:  20,
		Id:   1,
	}
	SetValue(&o)
	fmt.Println(o)
}

func SetValue(o interface{}) {
	// 判断是否为指针、是否能够修改
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && v.Elem().CanSet() {
		v = v.Elem()
	} else {
		panic("非指针或不可修改")
		return
	}

	if f := v.FieldByName("Name"); f.IsValid() && f.Kind() == reflect.String {
		f.SetString("hi")
	}
}
