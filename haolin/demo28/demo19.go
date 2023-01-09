package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("present time :%v\n", time.Now())
	expirationTime := <-timer.C
	fmt.Printf("Expiration time :%v.\n", expirationTime)
	fmt.Printf("Stop time:%v.\n", timer.Stop())
}
