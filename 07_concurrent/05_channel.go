package main

import (
	"fmt"
	"time"
)

var c = make(chan int)

func Print_7_5(word string) {
	for _, w := range word {
		fmt.Printf("%c", w)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func person1() {
	Print_7_5("hello")
	c <- 1
}
func person2() {
	<-c
	Print_7_5("world")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
