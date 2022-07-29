package main

import (
	"fmt"
	"os"
)

func main() {
	path := `./06_file/files/demo1.txt`
	err := os.Remove(path)
	if err != nil {
		fmt.Println("del failed, err: ", err)
		return
	}
	fmt.Println("del success")
}
