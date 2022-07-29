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
	// 获取文件路径及文件名
	fmt.Println("请输入要传输的文件路径:")
	buf := make([]byte, 1024)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	filePath := strings.TrimRight(string(buf[:n]), "\r\n")
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := fileInfo.Name()

	// 创建连接
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(conn)

	// 发送文件名
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 等待响应 ok
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 响应正确则开始传递文件
	if "ok" == string(buf[:n]) {
		// 获取文件
		var file *os.File
		file, err = os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}(file)
		// 传递
		for {
			n, err = file.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("文件传输完毕")
					return
				}
				fmt.Println(err)
				return
			}
			_, err = conn.Write(buf[:n])
			if err != nil {
				return
			}
		}
	} else {
		fmt.Println("验证失败")
		return
	}
}
