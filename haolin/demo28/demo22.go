package main

import (
	"fmt"
	"time"
)

func main() {
	var intChan = make(chan int, 1)
	var ticker = time.NewTicker(time.Second)
	stop := make(chan struct{}, 1)
	syncChan := make(chan struct{}, 1)
	go func() {
		//Loop:
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			case <-stop:
				ticker.Stop()

				//break Loop
			}
		}
		fmt.Println("End.[sender]")
		syncChan <- struct{}{}
	}()
	var sum int
	for e := range intChan {
		fmt.Printf("Received:%v\n", e)
		sum += e
		if sum > 10 {
			//ticker.Stop()
			stop <- struct{}{}
			fmt.Printf("Got:%v\n", sum)
			break
		}
	}
	fmt.Println("End.[receiver]")
	<-syncChan
}
