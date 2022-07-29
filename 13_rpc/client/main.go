package main

import (
	"fmt"
	"net/rpc"
)

type Params struct {
	A, B int
}

type Result struct {
	Chu, Yu int
}

func main() {
	http, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}

	params1 := &Params{
		A: 1,
		B: 2,
	}
	r1 := new(int)
	http.Call("Num.Add", params1, r1)
	fmt.Println(*r1)

	params2 := 5
	r2 := new(Result)
	http.Call("Num.Divide", params2, r2)
	fmt.Printf("%#v\n", r2)
}
