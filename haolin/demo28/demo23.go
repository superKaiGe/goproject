package main

import (
	"fmt"
	"time"
)

type Gen struct {
	timeoutNS time.Duration
	lps       int
}

func Test(gen Gen) {
	gen.timeoutNS = 50 * time.Millisecond
	gen.lps = 1000
	var total64 = int64(gen.timeoutNS)/int64(1e9/gen.lps) + 1
	fmt.Println(int64(gen.timeoutNS))
	fmt.Println(gen.timeoutNS)
	fmt.Println(1e9)
	//fmt.Println(gen.lps)
	fmt.Println(total64)
}
func main() {
	//Test(Gen{})
	var MaxInt32 = 1<<31 - 1
	fmt.Println(MaxInt32)
}
