package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个协程就加 1
		go func() {
			defer wg.Done() // 协程退出时自动减 1
			fmt.Println(i)
		}()
	}
	wg.Wait() // 等待所有协程执行完毕
}
