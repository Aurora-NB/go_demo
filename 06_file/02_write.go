package main

import (
	"fmt"
	"os"
)

func main() {
	path := `./06_file/files/demo.txt`
	WriteFile(path)
}

func WriteFile(path string) {
	// 新建文件
	fmt.Println(path)
	file, err := os.Create(path)
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

	// 准备写入内容
	var buf = ""
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		//写入文件
		_, err := file.WriteString(buf)
		if err != nil {
			fmt.Println("写文件错误")
			return
		}
	}

}
