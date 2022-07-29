package main

import (
	"fmt"
	"time"
)

func testStop() {
	timer := time.NewTimer(2 * time.Second)
	go func() {
		t := <-timer.C
		fmt.Println(t)
	}()
	timer.Stop()
	for {

	}
}
func testReset() {
	timer := time.NewTimer(5 * time.Second)
	timer.Reset(time.Second)
	go func() {
		t := <-timer.C
		fmt.Println(t)
	}()
	for {

	}
}
func main() {
	testReset()
}
