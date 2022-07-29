package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	str := "abc afq adc bbb aaa eee qqq"
	reg := regexp.MustCompile("a.c")
	res := reg.FindAllStringSubmatch(str, -1)
	fmt.Println(res)

	reg2 := regexp.MustCompile(`^(.*?)[+](.*?)i$`)
	res2 := reg2.FindStringSubmatch("10+-5i")
	for i, re := range res2 {
		fmt.Println(i, re)
	}
}

func getInfo(num string) (int, int) {
	reg2 := regexp.MustCompile(`^(.*?)[+](.*?)i$`)
	res := reg2.FindStringSubmatch(num)
	r1, _ := strconv.Atoi(res[1])
	r2, _ := strconv.Atoi(res[2])
	return r1, r2
}
