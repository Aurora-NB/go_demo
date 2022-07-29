package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("服务启动成功")

	for {
		// 阻塞等待连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 开启一个携程处理连接
		go HandlerConn(conn)
	}

}

func HandlerConn(conn net.Conn) {
	// 处理断开连接
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	buf := make([]byte, 1024)
	addr := conn.RemoteAddr().String()
	fmt.Printf("{%s}已连接\n", addr)

	// 开始读取信息
	for {
		n, err := conn.Read(buf)
		msg := string(buf[:n])
		if err != nil {
			fmt.Println("断开连接...")
			return
		}
		if "exit\r\n" == msg || "exit\n" == msg || "exit" == msg {
			fmt.Printf("{%s}已断开连接\n", addr)
			return
		}
		fmt.Printf("{%s}：%s\n", addr, msg)
		_, _ = conn.Write([]byte(strings.ToUpper(msg)))
	}
}
