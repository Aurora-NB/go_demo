package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get(`https://www.baidu.com`)
	if err != nil {
		return
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	buf := make([]byte, 1024)
	str := ""
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println(err)
			break
		}
		str += string(buf[:n])
	}
	fmt.Println(str)
}
