package main

import "fmt"

func main() {
	var m1 map[int]string
	fmt.Println(m1)

	m2 := make(map[int]string)
	m2[0] = "hello"
	fmt.Println(m2)

	m3 := make(map[int]string, 10)
	m3[2] = "world"
	fmt.Println(m3)

	m4 := map[int]string{1: "hello", 2: "world", 3: "!"}
	fmt.Println(m4)

	for k, v := range m4 {
		fmt.Printf("m4[%d] = %s\n", k, v)
	}

	delete(m4, 3)
	fmt.Println(m4)

	v1, exist1 := m4[1]
	v2, exist2 := m4[0]
	if exist1 {
		fmt.Println(v1)
	}
	if exist2 {
		fmt.Println(v2)
	}
}
