package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var mailbox uint8
	var lock sync.RWMutex
	//sendCond 代表专用于发信的条件变量
	sendCond := sync.NewCond(&lock)
	//recvCond 代表用于收信的条件变量
	recvCond := sync.NewCond(lock.RLocker())
	//sign 用于传递演示完成的信号
	sign := make(chan struct{}, 3)
	max := 50
	max1 := 20
	go func(max int) { //用于发信
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)
	go func(max1 int) { //用于收信
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max1; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", j)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max1)
	<-sign
	<-sign
}
