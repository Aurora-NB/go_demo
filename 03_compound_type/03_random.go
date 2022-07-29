package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		println(rand.Int())
	}

	for i := 0; i < 10; i++ {
		println(rand.Intn(50))
	}

}
