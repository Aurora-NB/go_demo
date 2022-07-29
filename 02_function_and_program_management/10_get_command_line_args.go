package main

import (
	"os"
)

func main() {
	list := os.Args
	for i, d := range list {
		println(i, d)
	}
}
