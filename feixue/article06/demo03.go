package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int32
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
func incCount() {
	defer wg.Done()
	for i := 0; i < 1050; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}
