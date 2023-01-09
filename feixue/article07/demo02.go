package main

import "fmt"

func main() {
	one := make(chan int)
	two := make(chan int)
	go func() {
		one <- 100
		one <- 200
	}()
	go func() {
		v := <-one
		two <- v
	}()
	fmt.Println(<-two)
}
