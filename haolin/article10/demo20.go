package main

import "fmt"

func main() {
	ch1 := make(chan int, 6)
	ch1 <- 2
	ch1 <- 10
	ch1 <- 3
	ch1 <- 3
	ch1 <- 3
	elem := <-ch1
	fmt.Printf("The first element of received from channel ch1: %v\n", elem)
}
