package main

import "fmt"

func main() {
	arr := [10]int{4, 7, 5, 9, 2, 1, 0, 3, 8, 6}
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println(arr)

}
