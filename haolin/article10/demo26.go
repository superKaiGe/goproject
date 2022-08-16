package main

import (
	"fmt"
	"time"
)

func main() {
	//timeout := 1
	TimeOut(time.Second * 1)
}
func TimeOut(timeout time.Duration) {
	ch_to := make(chan bool, 1)
	go func() {
		time.Sleep(timeout)
		ch_to <- true
	}()
	ch_do := make(chan int, 1)
	go func() {
		time.Sleep(3e9)
		ch_do <- 1
	}()
	select {
	case i := <-ch_do:
		fmt.Println("do something,id:", i)
	case <-ch_to:
		fmt.Println("timeout")
		break

	}
}
