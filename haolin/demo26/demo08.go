package main

import (
	"fmt"
	"time"
)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
}
func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 <-chan struct{}) {
	<-syncChan1
	fmt.Println("Receive a sync singal and wait a second...[receiver]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("Received:", elem, "[receiver]")
		} else {
			break
		}
	}
	fmt.Println("Stopped.[receiver]")
	//syncChan2 <- struct{}{}
}

//func send(strChan ) {
//
//}
