package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接到服务器
	dial, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("连接到：", dial.RemoteAddr(), "本地地址：", dial.LocalAddr())
	defer func(dial net.Conn) {
		err := dial.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(dial)

	// 发送消息
	msg := []byte("hello world")
	n, err := dial.Write(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("发送了%d个字节,内容为:%s\n", n, msg)

	// 提醒服务端断开连接
	_, err2 := dial.Write([]byte("000"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}

}
