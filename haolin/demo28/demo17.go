package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	fmt.Printf("aaa\n")
	close(intChan)
	syncChan := make(chan struct{}, 1)
	go func() {
		fmt.Printf("bbb\n")
	Loop:
		for {
			select {
			case e, ok := <-intChan:
				if !ok {
					fmt.Println("End.")
					break Loop
				}
				fmt.Printf("Received:%v\n", e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
}
