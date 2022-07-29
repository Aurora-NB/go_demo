package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func main() {
	// strings
	fmt.Println(strings.Contains("abcde", "abc"))
	fmt.Println(strings.Contains("abcde", "aaa"))
	fmt.Println(strings.Contains("abcde", ""))

	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, " "))

	fmt.Println(strings.Index("abc", "bc"))
	fmt.Println(strings.Index("abc", "cc"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println(strings.Replace("hello hello hello", "llo", "lp", 2))
	fmt.Println(strings.Replace("hello hello hello", "llo", "lp", -1))

	fmt.Println(strings.Split("hi,hello,good", ","))
	fmt.Println(strings.Split("hi,hello,good", ""))

	fmt.Println(strings.Trim("!!! hello world !!!", "! "))

	fmt.Println(strings.Fields("hello hello hello"))

	// Append
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10) //以10进制方式追加
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	str = strconv.AppendFloat(str, 3.14, 'f', -1, 64)

	fmt.Println(string(str)) //4567false"abcdefg"'单'3.14

	// Format
	a1 := strconv.FormatBool(false)
	b1 := strconv.FormatInt(1234, 10)
	c1 := strconv.FormatUint(12345, 10)
	d1 := strconv.Itoa(1023)

	fmt.Println(a1, b1, c1, d1) //false 1234 12345 1023

	// Parse
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e) //false 123.23 1234 12345 1023
}
