package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go test()
	for i := 0; i < 15; i++ {
		fmt.Println("main hello golang" + strconv.Itoa(i))
	}
	//time.Sleep(time.Second * 1)
}
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test hello world" + strconv.Itoa(i))
	}
	time.Sleep(time.Second * 3)
}
