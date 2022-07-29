package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	var err error = nil
	// 监听端口
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("开始等待文件传输...")

	// 创建连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 接收文件名
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := strings.TrimRight(string(buf[:n]), "\r\n")
	_, err = conn.Write([]byte("ok"))
	if err != nil {
		return
	}

	//	接收文件
	file, err := os.Create("./" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	// 接收文件
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件传输完毕")
				return
			}
			fmt.Println(err)
			return
		}
		_, err = file.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
