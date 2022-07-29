package main

import "fmt"

type person struct {
	name string
}

func main() {
	// 通过 if 进行断言
	arr := []interface{}{1, "hi", person{name: "ming"}}
	for index, data := range arr {
		if value, ok := data.(int); ok == true {
			fmt.Printf("%d 个参数为, %d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("%d 个参数为, %s\n", index, value)
		} else if value, ok := data.(person); ok == true {
			fmt.Printf("%d 个参数为, %+v\n", index, value)
		}
	}
	// 通过 switch 进行断言
	for index, data := range arr {
		switch value := data.(type) {
		case string:
			fmt.Printf("%d 个参数为, %s\n", index, value)
		case int:
			fmt.Printf("%d 个参数为, %d\n", index, value)
		case person:
			fmt.Printf("%d 个参数为, %+v\n", index, value)
		}
	}
}
