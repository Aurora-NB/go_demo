package main

import (
	"fmt"
	"net"
)

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
	addr := listener.Addr()
	fmt.Println("监听地址:", addr)

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
	fmt.Println("连接创建成功，本地地址：", conn.LocalAddr(), "远程地址：", conn.RemoteAddr())

	// 获取信息
	buf := make([]byte, 10*1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		msg := string(buf[:n])
		fmt.Println("接收到信息：", msg)
		if msg == "000\n" || msg == "000" {
			fmt.Println("通信结束")
			break
		}
	}
}
