package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Num int

type Params struct {
	A, B int
}
type Result struct {
	Chu, Yu int
}

func (s Num) Add(p *Params, res *int) error {
	*res = p.A + p.B
	return nil
}

func (s Num) Divide(d int, res *Result) error {
	if d == 0 {
		return errors.New("除数不能为0")
	}
	res.Chu = int(s) / d
	res.Yu = int(s) % d
	return nil
}

func main() {
	num := new(Num)
	*num = 9

	err := rpc.Register(num)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		err := http.Serve(listen, nil)
		if err != nil {
			fmt.Println(err)
		}
	}()
	select {}
}
