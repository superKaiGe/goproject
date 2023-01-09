package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch2 <- "ch2 value"
		ch1 <- "ch1 value"
	}()
	go func() {
		<-ch1
		<-ch2
	}()
	fmt.Println("main done")
}
