package main

import (
	"fmt"
	"reflect"
)

type People_10_1 struct {
	key int
}

type User_10_1 struct {
	People_10_1
	Name string
	Age  int
	Id   int
}

func (u User_10_1) Hello(s string) {
	fmt.Println(s, "hello")
}

func main() {
	o := User_10_1{
		Name: "king",
		Age:  20,
		Id:   1,
	}
	//getInfo1(&o)
	getInfo2(o)
}

func getInfo1(o interface{}) {
	// 获取类型
	t := reflect.TypeOf(o)

	if k := t.Kind(); k != reflect.Struct {
		panic("获取信息必须传入实例")
		return
	}

	fmt.Println("名称:", t.Name())
	fmt.Println("全名:", t.String())
	fmt.Println("字段数:", t.NumField())
	fmt.Println("方法数:", t.NumMethod())
	fmt.Println()

	// 获取各字段值
	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		// 获取字段
		f := t.Field(i)
		// 输出每个字段 名称、类型、值
		fmt.Printf("%-6s : %-6s : %v\n", f.Name, f.Type, v.Field(i).Interface())
	}
	fmt.Println()

	// 方法信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%-6s : %-6s", m.Name, m.Type)
	}
}

func getInfo2(o interface{}) {
	t := reflect.TypeOf(o)
	// 匿名字段的 Anonymous 是 true
	// 此处是获取第一个属性的第一个属性
	fmt.Println(t.FieldByIndex([]int{0, 0}))
}
