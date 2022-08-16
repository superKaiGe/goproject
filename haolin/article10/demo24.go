package main

import "fmt"

func main() {
	//ch 用来同步两个协程交替执行
	ch := make(chan int)
	//ch_end 用来阻塞主程序 让两个协程可以执行完
	ch_end := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- 1
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		ch_end <- 1
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			<-ch
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()
	<-ch_end
}
