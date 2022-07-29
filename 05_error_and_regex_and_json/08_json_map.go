package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{})
	m["name"] = "hi"
	m["number"] = 3
	m["isOk"] = true
	m["dep"] = []string{"a", "b", "c"}
	m["price"] = 3.14
	m["map"] = map[string]interface{}{"test": "test"}
	//j, err := json.Marshal(comp)
	j, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))

	var m2 map[string]interface{}
	err2 := json.Unmarshal(j, &m2)
	if err2 != nil {
		return
	}
	fmt.Println(m2)

	// 通过解析字符串得到的 map 必须使用类型断言才能得到原来的类型，如果内部有数组还需要在此类型断言
	for _, v := range m2 {
		switch data := v.(type) {
		case string:
			fmt.Println(data)
		case int:
			fmt.Println(data)
		}
		// ...
	}
}
