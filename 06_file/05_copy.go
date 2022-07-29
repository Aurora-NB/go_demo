package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("应包含2个参数")
		return
	}
	fromPath, toPath := args[1], args[2]
	if fromPath == toPath {
		fmt.Println("路径不能相同")
		return
	}
	file, err := os.Open(fromPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	file2, err2 := os.Create(toPath)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file2)
	buf := make([]byte, 1024*4)
	for {
		n, err3 := file.Read(buf)

		if err3 != nil {
			if err3 == io.EOF {
				break
			}
			fmt.Println(err3)
			return
		}
		buf = buf[:n]
		_, err4 := file2.Write(buf)
		if err4 != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("写入完成")
}
