package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try to lock for reading...[%d]\n", i)
			rwm.RLock()
			fmt.Printf("locked for reading.[%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("try to unlock for reading...[%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("unlocked for reading.[%d]\n", i)
		}(i)
	}
	time.Sleep(time.Microsecond * 100)
	fmt.Println("try to lock for reading...")
	rwm.Lock()
	fmt.Println("locked for writing.")
}
