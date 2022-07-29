package main

import (
	"fmt"
	"net"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(listener)

	// 创建连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	// 获取信息
	buf := make([]byte, 10*1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := string(buf[:n])
	fmt.Println(msg)

}
