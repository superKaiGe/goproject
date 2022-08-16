package main

func main() {
	//示例1
	//ch1 := make(chan int, 1)
	//ch1 <- 1
	//ch1 <- 1 //通道已满 造成阻塞
	//ch2 := make(chan int, 1)
	//elem, ok := <-ch2 //通道已空,造成阻塞
	//_, _ = elem, ok
	//
	//var ch3 chan int
	//ch3 <- 1 //通道值为nil 造成永久的阻塞
	////<-ch3//通道里值为nil 造成永久的阻塞
	//_ = ch3
	//var chan1 chan int
	//<-chan1
	////chan1 <- 13
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

}
