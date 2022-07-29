package main

import (
	"fmt"
	"time"
)

func testAfter() {
	timer := time.After(2 * time.Second)
	<-timer
	fmt.Println("时间到")
}

func testTimer() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now())
	t := <-timer.C
	fmt.Println(t)
}

func main() {
	testTimer()
	//testAfter()
}
