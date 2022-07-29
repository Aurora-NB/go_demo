package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	ip := strings.Split((conn.LocalAddr().(*net.UDPAddr)).IP.String(), ":")[0]
	fmt.Println(ip)
}
