package main

import (
	"fmt"
	"time"
)

//Counter 代表计算器的类型
type Counter struct {
	count int
}

var mapChan = make(chan map[string]*Counter, 1)

func main() {
	synChan := make(chan struct{}, 2)
	go func() { //用于接收操作
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped.[receiver]")
		synChan <- struct{}{}
	}()
	go func() { //用于演示发送操作
		countMap := map[string]*Counter{"count": &Counter{}}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map:%v.[sender]\n", countMap)
		}
		close(mapChan)
		synChan <- struct{}{}
	}()
	<-synChan
	<-synChan
}
func (c *Counter) String() string {
	return fmt.Sprintf("{count:%d}", c.count)
}
