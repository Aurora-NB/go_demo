package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	host := "8.136.225.205"
	port := "8888"
	// 创建连接
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("已连接到聊天室...")
	// 监听服务端消息
	go func() {
		buf := make([]byte, 5*1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				return
			}
			msg := string(buf[:n])
			fmt.Print(msg)
		}
	}()

	// 向服务端发送信息
	buf := make([]byte, 5*1024)
	for {
		var msg string
		n, _ := os.Stdin.Read(buf)
		msg = string(buf[:n])
		msg = strings.Trim(msg, "\r\n")
		if "func|exit" == msg {
			fmt.Printf("close connect...\n")
			return
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("write failed, err: %v\n", err)
		}
	}
}
