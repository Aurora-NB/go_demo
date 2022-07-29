package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := `./06_file/files/demo.txt`
	//ReadFile(path)
	ReadLine(path)
}

func ReadFile(path string) {
	// 打开文件
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭错误")
		}
	}(file)

	// 读取
	buf := make([]byte, 1024)
	n, err2 := file.Read(buf)
	if err2 != nil && err2 != io.EOF { // 读取错误，且没读完
		return
	}
	buf = buf[:n]
	fmt.Println(string(buf))
}

func ReadLine(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		buf, err2 := reader.ReadBytes('\n')
		//buf, _, err2 := reader.ReadLine()
		if err2 != nil {
			if err2 == io.EOF {
				break
			}
			fmt.Println(err2)
		}
		fmt.Printf("%s", string(buf))
	}
}
