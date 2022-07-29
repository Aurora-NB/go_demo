package main

import "fmt"

func main() {
	str := "abc"
	for i, c := range str {
		fmt.Printf("str[%d] = %c\n", i, c)
	}

	num := 0
	for {
		num++
		fmt.Println(num)
		if num == 10 {
			break
		}
	}
}
