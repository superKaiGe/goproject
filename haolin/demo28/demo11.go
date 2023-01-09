package main

import "fmt"

//var ch chan int
var ch = make(chan int, 1)

func main() {
	ch := make(chan int, 1)
	//从一个还未被初始化的通道里接收元素值 会导致当前goroutine的永久阻塞 使用for语句也不例外
	for i2 := range ch {
		fmt.Printf("%s\n", i2)
	}
}
