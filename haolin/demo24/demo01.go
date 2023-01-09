package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "send"
	}()
	fmt.Println(<-ch)
}
