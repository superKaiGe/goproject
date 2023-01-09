package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var mailbox uint8
	var lock sync.Mutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(&lock)
	//send 代表用于发信的函数
	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		log.Printf("sender [%d-%d]: the mailbox is empty.",
			id, index)
		mailbox = 1
		log.Printf("sender [%d-%d]: the letter has been sent.",
			id, index)
		lock.Unlock()
		recvCond.Broadcast()
	}
	//recv代表用于收信的函数
	recv := func(id, index int) {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		log.Printf("receiver [%d-%d]: the mailbox is full.",
			id, index)
		mailbox = 0
		log.Printf("receiver [%d-%d]: the letter has been received.",
			id, index)
		lock.Unlock()
		sendCond.Signal()
	}
	//sign 用于传递演示完成的信号
	sign := make(chan struct{}, 3)
	max := 6
	go func(id, max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			send(id, i)
		}
	}(0, max)
	go func(id, max int) { //用于收信
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, j)
		}
	}(1, max/2)
	go func(id, max int) { //用于收信
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= max; k++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, k)
		}
	}(2, max/2)
	<-sign
	<-sign
	<-sign
}
