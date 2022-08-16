package main

import "fmt"

func main() {
	ch1 := make(chan int)
	//ch1 <- 1
	//<-ch1
	go func() {
		a := <-ch1
		fmt.Println(a)
	}()
	ch1 <- 1

	fmt.Println("------------")

	//ch := make(chan int, 1)
	//ch <- 1
	//fmt.Println(<-ch)
}
