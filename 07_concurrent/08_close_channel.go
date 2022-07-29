package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// range 自动判断管道是否关闭
	for num := range ch {
		fmt.Println(num)
	}

	// 通过 ok 判断管道是否关闭
	//for {
	//	if i, ok := <-ch; ok == true {
	//		fmt.Println(i)
	//	} else {
	//		break
	//	}
	//}
}
