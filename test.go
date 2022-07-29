package main

import (
	"fmt"
	"unicode"
)

type Student struct {
	id *int
}

func main() {
	i := 1
	p := &i
	a := Student{p}
	b := a
	var m = map[int]string{}
	fmt.Println(m)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}

func reverseOnlyLetters(s string) string {
	b := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		for !unicode.IsLetter(rune(b[i])) {
			i++
		}
		for !unicode.IsLetter(rune(b[i])) {
			j--
		}
		if i >= j {
			break
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
