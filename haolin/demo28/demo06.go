package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped.[receiver]")
		syncChan <- struct{}{}
	}()
	go func() { //用于演示发送操作
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			fmt.Printf("%v\n", countMap)
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("%v\n", countMap)
			fmt.Printf("The count map:%v.[sender]\n", countMap)

		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
