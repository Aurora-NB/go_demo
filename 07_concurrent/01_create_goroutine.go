package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			fmt.Println("go: ", i)
			i++
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		fmt.Println("main: ", i)
		i++
		time.Sleep(time.Second)
	}
}
