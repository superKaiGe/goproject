package main

import (
	"runtime"
)

func main() {
	go println("aaa")
	//time.Sleep(time.Millisecond)
	runtime.Gosched()
}
