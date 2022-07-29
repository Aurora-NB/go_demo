package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("已连接到", conn.RemoteAddr())
	// 监听服务端消息
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				return
			}
			msg := string(buf[:n])
			fmt.Println(msg)
		}
	}()

	// 向服务端发送信息
	for {
		var msg string
		fmt.Scanf("%s", &msg)
		if "exit\r\n" == msg || "exit\n" == msg || "exit" == msg {
			fmt.Printf("{%s}已断开连接\n")
			return
		}
		_, _ = conn.Write([]byte(msg))
	}
}
